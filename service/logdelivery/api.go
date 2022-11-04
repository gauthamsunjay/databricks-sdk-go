// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package logdelivery

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewLogDelivery(client *client.DatabricksClient) LogDeliveryService {
	return &LogDeliveryAPI{
		client: client,
	}
}

type LogDeliveryAPI struct {
	client *client.DatabricksClient
}

// Create a new log delivery configuration
//
// Creates a new Databricks log delivery configuration to enable delivery of the
// specified type of logs to your storage location. This requires that you
// already created a [credential object](#operation/create-credential-config)
// (which encapsulates a cross-account service IAM role) and a [storage
// configuration object](#operation/create-storage-config) (which encapsulates
// an S3 bucket).
//
// For full details, including the required IAM role policies and bucket
// policies, see [Deliver and access billable usage
// logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html)
// or [Configure audit
// logging](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
//
// **Note**: There is a limit on the number of log delivery configurations
// available per account (each limit applies separately to each log type
// including billable usage and audit logs). You can create a maximum of two
// enabled account-level delivery configurations (configurations without a
// workspace filter) per type. Additionally, you can create two enabled
// workspace-level delivery configurations per workspace for each log type,
// which means that the same workspace ID can occur in the workspace filter for
// no more than two delivery configurations per log type.
//
// You cannot delete a log delivery configuration, but you can disable it (see
// [Enable or disable log delivery
// configuration](#operation/patch-log-delivery-config-status)).
func (a *LogDeliveryAPI) CreateLogDeliveryConfig(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", request.AccountId)
	err := a.client.Post(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Get all log delivery configurations
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
//
// Use GetAllLogDeliveryConfigsAll() to get all LogDeliveryConfiguration instances
func (a *LogDeliveryAPI) GetAllLogDeliveryConfigs(ctx context.Context, request GetAllLogDeliveryConfigsRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", request.AccountId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

// GetAllLogDeliveryConfigsAll returns all LogDeliveryConfiguration instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LogDeliveryAPI) GetAllLogDeliveryConfigsAll(ctx context.Context, request GetAllLogDeliveryConfigsRequest) ([]LogDeliveryConfiguration, error) {
	response, err := a.GetAllLogDeliveryConfigs(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.LogDeliveryConfigurations, nil
}

// Get all log delivery configurations
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *LogDeliveryAPI) GetAllLogDeliveryConfigsByAccountId(ctx context.Context, accountId string) (*WrappedLogDeliveryConfigurations, error) {
	return a.GetAllLogDeliveryConfigs(ctx, GetAllLogDeliveryConfigsRequest{
		AccountId: accountId,
	})
}

// Get log delivery configuration
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) GetLogDeliveryConfig(ctx context.Context, request GetLogDeliveryConfigRequest) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", request.AccountId, request.LogDeliveryConfigurationId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Get log delivery configuration
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) GetLogDeliveryConfigByAccountIdAndLogDeliveryConfigurationId(ctx context.Context, accountId string, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error) {
	return a.GetLogDeliveryConfig(ctx, GetLogDeliveryConfigRequest{
		AccountId:                  accountId,
		LogDeliveryConfigurationId: logDeliveryConfigurationId,
	})
}

// Enable or disable log delivery configuration
//
// Enables or disables a log delivery configuration. Deletion of delivery
// configurations is not supported, so disable log delivery configurations that
// are no longer needed. Note that you can't re-enable a delivery configuration
// if this would violate the delivery configuration limits described under
// [Create log delivery](#operation/create-log-delivery-config).
func (a *LogDeliveryAPI) PatchLogDeliveryConfigStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", request.AccountId, request.LogDeliveryConfigurationId)
	err := a.client.Patch(ctx, path, request)
	return err
}
