package compute

import (
	"context"
	"fmt"
	"sync"

	"github.com/databricks/databricks-sdk-go/logger"
)

// getOrCreateClusterMutex guards "mounting" cluster creation to prevent multiple
// redundant instances created at the same name. Compute package private property.
// https://github.com/databricks/terraform-provider-databricks/issues/445
var getOrCreateClusterMutex sync.Mutex

func (ci *ClusterInfo) IsRunningOrResizing() bool {
	return ci.State == StateRunning || ci.State == StateResizing
}

// use mutex around starting a cluster by ID
var mu sync.Mutex

func (a *ClustersAPI) EnsureClusterIsRunning(ctx context.Context, clusterId string) error {
	mu.Lock()
	defer mu.Unlock()
	info, err := a.GetByClusterId(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("get cluster info: %w", err)
	}
	switch info.State {
	case StateRunning:
		return nil
	case StateTerminated:
		// TODO: add StateTerminating: self.wait_get_cluster_terminated(cluster_id) & self.start(cluster_id).result()
		_, err = a.StartByClusterIdAndWait(ctx, clusterId)
		if err != nil {
			return fmt.Errorf("start: %w", err)
		}
		return nil
	case StatePending, StateResizing, StateRestarting:
		_, err = a.GetByClusterIdAndWait(ctx, clusterId)
		if err != nil {
			return fmt.Errorf("wait: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("cluster %s is in %s state: %s", info.ClusterName, info.State, info.StateMessage)
	}
}

// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
func (a *ClustersAPI) GetOrCreateRunningCluster(ctx context.Context, name string, custom ...CreateCluster) (c *ClusterInfo, err error) {
	getOrCreateClusterMutex.Lock()
	defer getOrCreateClusterMutex.Unlock()
	if len(custom) > 1 {
		err = fmt.Errorf("you can only specify 1 custom cluster conf, not %d", len(custom))
		return
	}
	clusters, err := a.ListAll(ctx, ListClustersRequest{})
	if err != nil {
		return
	}
	for _, cl := range clusters {
		if cl.ClusterName != name {
			continue
		}
		logger.Infof(ctx, "Found reusable cluster '%s'", name)
		if cl.IsRunningOrResizing() {
			return &cl, nil
		}
		started, err := a.StartByClusterIdAndWait(ctx, cl.ClusterId)
		if err != nil {
			logger.Infof(ctx, "Cluster %s cannot be started, creating an autoterminating cluster: %s", name, err)
			break
		}
		return started, nil
	}
	nodeTypes, err := a.ListNodeTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list node types: %w", err)
	}
	smallestNodeType, err := nodeTypes.Smallest(NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		return nil, err
	}
	logger.Infof(ctx, "Creating an autoterminating cluster with node type %s", smallestNodeType)
	versions, err := a.SparkVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("list spark versions: %w", err)
	}
	version, err := versions.Select(SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		return nil, err
	}
	r := CreateCluster{
		NumWorkers:             1,
		ClusterName:            name,
		SparkVersion:           version,
		NodeTypeId:             smallestNodeType,
		AutoterminationMinutes: 10,
	}
	api, ok := a.impl.(*clustersImpl)
	if !ok {
		return nil, fmt.Errorf("cannot get raw clusters API")
	}
	if api.client.Config.IsAws() {
		r.AwsAttributes = &AwsAttributes{
			Availability: "SPOT",
		}
	}
	if len(custom) == 1 {
		r = custom[0]
	}
	return a.CreateAndWait(ctx, r)
}
