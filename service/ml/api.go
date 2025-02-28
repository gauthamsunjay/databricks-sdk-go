// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Experiments, Model Registry, etc.
package ml

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewExperiments(client *client.DatabricksClient) *ExperimentsAPI {
	return &ExperimentsAPI{
		impl: &experimentsImpl{
			client: client,
		},
	}
}

// Experiments are the primary unit of organization in MLflow; all MLflow runs
// belong to an experiment. Each experiment lets you visualize, search, and
// compare runs, as well as download run artifacts or metadata for analysis in
// other tools. Experiments are maintained in a Databricks hosted MLflow
// tracking server.
//
// Experiments are located in the workspace file tree. You manage experiments
// using the same tools you use to manage other workspace objects such as
// folders, notebooks, and libraries.
type ExperimentsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ExperimentsService)
	impl ExperimentsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ExperimentsAPI) WithImpl(impl ExperimentsService) *ExperimentsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Experiments API implementation
func (a *ExperimentsAPI) Impl() ExperimentsService {
	return a.impl
}

// Create experiment.
//
// Creates an experiment with a name. Returns the ID of the newly created
// experiment. Validates that another experiment with the same name does not
// already exist and fails if another experiment with the same name already
// exists.
//
// Throws `RESOURCE_ALREADY_EXISTS` if a experiment with the given name exists.
func (a *ExperimentsAPI) CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	return a.impl.CreateExperiment(ctx, request)
}

// Create a run.
//
// Creates a new run within an experiment. A run is usually a single execution
// of a machine learning or data ETL pipeline. MLflow uses runs to track the
// `mlflowParam`, `mlflowMetric` and `mlflowRunTag` associated with a single
// execution.
func (a *ExperimentsAPI) CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	return a.impl.CreateRun(ctx, request)
}

// Delete an experiment.
//
// Marks an experiment and associated metadata, runs, metrics, params, and tags
// for deletion. If the experiment uses FileStore, artifacts associated with
// experiment are also deleted.
func (a *ExperimentsAPI) DeleteExperiment(ctx context.Context, request DeleteExperiment) error {
	return a.impl.DeleteExperiment(ctx, request)
}

// Delete a run.
//
// Marks a run for deletion.
func (a *ExperimentsAPI) DeleteRun(ctx context.Context, request DeleteRun) error {
	return a.impl.DeleteRun(ctx, request)
}

// Delete a tag.
//
// Deletes a tag on a run. Tags are run metadata that can be updated during a
// run and after a run completes.
func (a *ExperimentsAPI) DeleteTag(ctx context.Context, request DeleteTag) error {
	return a.impl.DeleteTag(ctx, request)
}

// Get metadata.
//
// Gets metadata for an experiment.
//
// This endpoint will return deleted experiments, but prefers the active
// experiment if an active and deleted experiment share the same name. If
// multiple deleted experiments share the same name, the API will return one of
// them.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no experiment with the specified name
// exists.
func (a *ExperimentsAPI) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {
	return a.impl.GetByName(ctx, request)
}

// Get an experiment.
//
// Gets metadata for an experiment. This method works on deleted experiments.
func (a *ExperimentsAPI) GetExperiment(ctx context.Context, request GetExperimentRequest) (*Experiment, error) {
	return a.impl.GetExperiment(ctx, request)
}

// Get history of a given metric within a run.
//
// Gets a list of all values for the specified metric for a given run.
func (a *ExperimentsAPI) GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	return a.impl.GetHistory(ctx, request)
}

// Get a run.
//
// Gets the metadata, metrics, params, and tags for a run. In the case where
// multiple metrics with the same key are logged for a run, return only the
// value with the latest timestamp.
//
// If there are multiple values with the latest timestamp, return the maximum of
// these values.
func (a *ExperimentsAPI) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	return a.impl.GetRun(ctx, request)
}

// Get all artifacts.
//
// List artifacts for a run. Takes an optional `artifact_path` prefix. If it is
// specified, the response contains only artifacts with the specified prefix.",
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExperimentsAPI) ListArtifactsAll(ctx context.Context, request ListArtifactsRequest) ([]FileInfo, error) {
	var results []FileInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListArtifacts(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Files) == 0 {
			break
		}
		for _, v := range response.Files {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// List experiments.
//
// Gets a list of all experiments.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExperimentsAPI) ListExperimentsAll(ctx context.Context, request ListExperimentsRequest) ([]Experiment, error) {
	var results []Experiment
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListExperiments(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Experiments) == 0 {
			break
		}
		for _, v := range response.Experiments {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Log a batch.
//
// Logs a batch of metrics, params, and tags for a run. If any data failed to be
// persisted, the server will respond with an error (non-200 status code).
//
// In case of error (due to internal server error or an invalid request),
// partial data may be written.
//
// You can write metrics, params, and tags in interleaving fashion, but within a
// given entity type are guaranteed to follow the order specified in the request
// body.
//
// The overwrite behavior for metrics, params, and tags is as follows:
//
// * Metrics: metric values are never overwritten. Logging a metric (key, value,
// timestamp) appends to the set of values for the metric with the provided key.
//
// * Tags: tag values can be overwritten by successive writes to the same tag
// key. That is, if multiple tag values with the same key are provided in the
// same API request, the last-provided tag value is written. Logging the same
// tag (key, value) is permitted. Specifically, logging a tag is idempotent.
//
// * Parameters: once written, param values cannot be changed (attempting to
// overwrite a param value will result in an error). However, logging the same
// param (key, value) is permitted. Specifically, logging a param is idempotent.
//
// Request Limits ------------------------------- A single JSON-serialized API
// request may be up to 1 MB in size and contain:
//
// * No more than 1000 metrics, params, and tags in total * Up to 1000 metrics *
// Up to 100 params * Up to 100 tags
//
// For example, a valid request might contain 900 metrics, 50 params, and 50
// tags, but logging 900 metrics, 50 params, and 51 tags is invalid.
//
// The following limits also apply to metric, param, and tag keys and values:
//
// * Metric keyes, param keys, and tag keys can be up to 250 characters in
// length * Parameter and tag values can be up to 250 characters in length
func (a *ExperimentsAPI) LogBatch(ctx context.Context, request LogBatch) error {
	return a.impl.LogBatch(ctx, request)
}

// Log inputs to a run.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (a *ExperimentsAPI) LogInputs(ctx context.Context, request LogInputs) error {
	return a.impl.LogInputs(ctx, request)
}

// Log a metric.
//
// Logs a metric for a run. A metric is a key-value pair (string key, float
// value) with an associated timestamp. Examples include the various metrics
// that represent ML model accuracy. A metric can be logged multiple times.
func (a *ExperimentsAPI) LogMetric(ctx context.Context, request LogMetric) error {
	return a.impl.LogMetric(ctx, request)
}

// Log a model.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (a *ExperimentsAPI) LogModel(ctx context.Context, request LogModel) error {
	return a.impl.LogModel(ctx, request)
}

// Log a param.
//
// Logs a param used for a run. A param is a key-value pair (string key, string
// value). Examples include hyperparameters used for ML model training and
// constant dates and values used in an ETL pipeline. A param can be logged only
// once for a run.
func (a *ExperimentsAPI) LogParam(ctx context.Context, request LogParam) error {
	return a.impl.LogParam(ctx, request)
}

// Restores an experiment.
//
// Restore an experiment marked for deletion. This also restores associated
// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
// underlying artifacts associated with experiment are also restored.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if experiment was never created or was
// permanently deleted.
func (a *ExperimentsAPI) RestoreExperiment(ctx context.Context, request RestoreExperiment) error {
	return a.impl.RestoreExperiment(ctx, request)
}

// Restore a run.
//
// Restores a deleted run.
func (a *ExperimentsAPI) RestoreRun(ctx context.Context, request RestoreRun) error {
	return a.impl.RestoreRun(ctx, request)
}

// Search experiments.
//
// Searches for experiments that satisfy specified search criteria.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExperimentsAPI) SearchExperimentsAll(ctx context.Context, request SearchExperiments) ([]Experiment, error) {
	var results []Experiment
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.SearchExperiments(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Experiments) == 0 {
			break
		}
		for _, v := range response.Experiments {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Search for runs.
//
// Searches for runs that satisfy expressions.
//
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.",
//
// This method is generated by Databricks SDK Code Generator.
func (a *ExperimentsAPI) SearchRunsAll(ctx context.Context, request SearchRuns) ([]Run, error) {
	var results []Run
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.SearchRuns(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Runs) == 0 {
			break
		}
		for _, v := range response.Runs {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Set a tag.
//
// Sets a tag on an experiment. Experiment tags are metadata that can be
// updated.
func (a *ExperimentsAPI) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	return a.impl.SetExperimentTag(ctx, request)
}

// Set a tag.
//
// Sets a tag on a run. Tags are run metadata that can be updated during a run
// and after a run completes.
func (a *ExperimentsAPI) SetTag(ctx context.Context, request SetTag) error {
	return a.impl.SetTag(ctx, request)
}

// Update an experiment.
//
// Updates experiment metadata.
func (a *ExperimentsAPI) UpdateExperiment(ctx context.Context, request UpdateExperiment) error {
	return a.impl.UpdateExperiment(ctx, request)
}

// Update a run.
//
// Updates run metadata.
func (a *ExperimentsAPI) UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	return a.impl.UpdateRun(ctx, request)
}

func NewModelRegistry(client *client.DatabricksClient) *ModelRegistryAPI {
	return &ModelRegistryAPI{
		impl: &modelRegistryImpl{
			client: client,
		},
	}
}

// MLflow Model Registry is a centralized model repository and a UI and set of
// APIs that enable you to manage the full lifecycle of MLflow Models.
type ModelRegistryAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ModelRegistryService)
	impl ModelRegistryService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ModelRegistryAPI) WithImpl(impl ModelRegistryService) *ModelRegistryAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ModelRegistry API implementation
func (a *ModelRegistryAPI) Impl() ModelRegistryService {
	return a.impl
}

// Approve transition request.
//
// Approves a model version stage transition request.
func (a *ModelRegistryAPI) ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequest) (*ApproveTransitionRequestResponse, error) {
	return a.impl.ApproveTransitionRequest(ctx, request)
}

// Post a comment.
//
// Posts a comment on a model version. A comment can be submitted either by a
// user or programmatically to display relevant information about the model. For
// example, test results or deployment errors.
func (a *ModelRegistryAPI) CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error) {
	return a.impl.CreateComment(ctx, request)
}

// Create a model.
//
// Creates a new registered model with the name specified in the request body.
//
// Throws `RESOURCE_ALREADY_EXISTS` if a registered model with the given name
// exists.
func (a *ModelRegistryAPI) CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error) {
	return a.impl.CreateModel(ctx, request)
}

// Create a model version.
//
// Creates a model version.
func (a *ModelRegistryAPI) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	return a.impl.CreateModelVersion(ctx, request)
}

// Make a transition request.
//
// Creates a model version stage transition request.
func (a *ModelRegistryAPI) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error) {
	return a.impl.CreateTransitionRequest(ctx, request)
}

// Create a webhook.
//
// **NOTE**: This endpoint is in Public Preview.
//
// Creates a registry webhook.
func (a *ModelRegistryAPI) CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error) {
	return a.impl.CreateWebhook(ctx, request)
}

// Delete a comment.
//
// Deletes a comment on a model version.
func (a *ModelRegistryAPI) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {
	return a.impl.DeleteComment(ctx, request)
}

// Delete a model.
//
// Deletes a registered model.
func (a *ModelRegistryAPI) DeleteModel(ctx context.Context, request DeleteModelRequest) error {
	return a.impl.DeleteModel(ctx, request)
}

// Delete a model tag.
//
// Deletes the tag for a registered model.
func (a *ModelRegistryAPI) DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error {
	return a.impl.DeleteModelTag(ctx, request)
}

// Delete a model version.
//
// Deletes a model version.
func (a *ModelRegistryAPI) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {
	return a.impl.DeleteModelVersion(ctx, request)
}

// Delete a model version tag.
//
// Deletes a model version tag.
func (a *ModelRegistryAPI) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	return a.impl.DeleteModelVersionTag(ctx, request)
}

// Delete a transition request.
//
// Cancels a model version stage transition request.
func (a *ModelRegistryAPI) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) error {
	return a.impl.DeleteTransitionRequest(ctx, request)
}

// Delete a webhook.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Deletes a registry webhook.
func (a *ModelRegistryAPI) DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error {
	return a.impl.DeleteWebhook(ctx, request)
}

// Get the latest version.
//
// Gets the latest version of a registered model.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelRegistryAPI) GetLatestVersionsAll(ctx context.Context, request GetLatestVersionsRequest) ([]ModelVersion, error) {
	response, err := a.impl.GetLatestVersions(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.ModelVersions, nil
}

// Get model.
//
// Get the details of a model. This is a Databricks workspace version of the
// [MLflow endpoint] that also returns the model's Databricks workspace ID and
// the permission level of the requesting user on the model.
//
// [MLflow endpoint]: https://www.mlflow.org/docs/latest/rest-api.html#get-registeredmodel
func (a *ModelRegistryAPI) GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error) {
	return a.impl.GetModel(ctx, request)
}

// Get a model version.
//
// Get a model version.
func (a *ModelRegistryAPI) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	return a.impl.GetModelVersion(ctx, request)
}

// Get a model version URI.
//
// Gets a URI to download the model version.
func (a *ModelRegistryAPI) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	return a.impl.GetModelVersionDownloadUri(ctx, request)
}

// List models.
//
// Lists all available registered models, up to the limit specified in
// __max_results__.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelRegistryAPI) ListModelsAll(ctx context.Context, request ListModelsRequest) ([]Model, error) {
	var results []Model
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListModels(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.RegisteredModels) == 0 {
			break
		}
		for _, v := range response.RegisteredModels {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// List transition requests.
//
// Gets a list of all open stage transition requests for the model version.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelRegistryAPI) ListTransitionRequestsAll(ctx context.Context, request ListTransitionRequestsRequest) ([]Activity, error) {
	response, err := a.impl.ListTransitionRequests(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Requests, nil
}

// List registry webhooks.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Lists all registry webhooks.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelRegistryAPI) ListWebhooksAll(ctx context.Context, request ListWebhooksRequest) ([]RegistryWebhook, error) {
	var results []RegistryWebhook
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListWebhooks(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Webhooks) == 0 {
			break
		}
		for _, v := range response.Webhooks {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Reject a transition request.
//
// Rejects a model version stage transition request.
func (a *ModelRegistryAPI) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error) {
	return a.impl.RejectTransitionRequest(ctx, request)
}

// Rename a model.
//
// Renames a registered model.
func (a *ModelRegistryAPI) RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error) {
	return a.impl.RenameModel(ctx, request)
}

// Searches model versions.
//
// Searches for specific model versions based on the supplied __filter__.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelRegistryAPI) SearchModelVersionsAll(ctx context.Context, request SearchModelVersionsRequest) ([]ModelVersion, error) {
	var results []ModelVersion
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.SearchModelVersions(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.ModelVersions) == 0 {
			break
		}
		for _, v := range response.ModelVersions {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Search models.
//
// Search for registered models based on the specified __filter__.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ModelRegistryAPI) SearchModelsAll(ctx context.Context, request SearchModelsRequest) ([]Model, error) {
	var results []Model
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.SearchModels(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.RegisteredModels) == 0 {
			break
		}
		for _, v := range response.RegisteredModels {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// Set a tag.
//
// Sets a tag on a registered model.
func (a *ModelRegistryAPI) SetModelTag(ctx context.Context, request SetModelTagRequest) error {
	return a.impl.SetModelTag(ctx, request)
}

// Set a version tag.
//
// Sets a model version tag.
func (a *ModelRegistryAPI) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {
	return a.impl.SetModelVersionTag(ctx, request)
}

// Test a webhook.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Tests a registry webhook.
func (a *ModelRegistryAPI) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	return a.impl.TestRegistryWebhook(ctx, request)
}

// Transition a stage.
//
// Transition a model version's stage. This is a Databricks workspace version of
// the [MLflow endpoint] that also accepts a comment associated with the
// transition to be recorded.",
//
// [MLflow endpoint]: https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage
func (a *ModelRegistryAPI) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	return a.impl.TransitionStage(ctx, request)
}

// Update a comment.
//
// Post an edit to a comment on a model version.
func (a *ModelRegistryAPI) UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error) {
	return a.impl.UpdateComment(ctx, request)
}

// Update model.
//
// Updates a registered model.
func (a *ModelRegistryAPI) UpdateModel(ctx context.Context, request UpdateModelRequest) error {
	return a.impl.UpdateModel(ctx, request)
}

// Update model version.
//
// Updates the model version.
func (a *ModelRegistryAPI) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) error {
	return a.impl.UpdateModelVersion(ctx, request)
}

// Update a webhook.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Updates a registry webhook.
func (a *ModelRegistryAPI) UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) error {
	return a.impl.UpdateWebhook(ctx, request)
}
