package aihc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// BatchStopTrainingTasksV2
//
// PARAMS:
//   - request: the arguments to BatchStopTrainingTasksV2
//
// RETURNS:
//   - BatchStopTrainingTasksV2Response: The return type of the BatchStopTrainingTasksV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchStopTrainingTasksV2(request *BatchStopTrainingTasksV2Request) (*BatchStopTrainingTasksV2Response, error) {
	result := &BatchStopTrainingTasksV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchStopTrainingTasksV2Uri()).
		WithQueryParamFilter("action", "BatchStopJobs").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDataset
//
// PARAMS:
//   - request: the arguments to CreateDataset
//
// RETURNS:
//   - CreateDatasetResponse: The return type of the CreateDataset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDataset(request *CreateDatasetRequest) (*CreateDatasetResponse, error) {
	result := &CreateDatasetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDatasetUri()).
		WithQueryParamFilter("action", "CreateDataset").
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDatasetVersion
//
// PARAMS:
//   - request: the arguments to CreateDatasetVersion
//
// RETURNS:
//   - CreateDatasetVersionResponse: The return type of the CreateDatasetVersion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDatasetVersion(request *CreateDatasetVersionRequest) (*CreateDatasetVersionResponse, error) {
	result := &CreateDatasetVersionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDatasetVersionUri()).
		WithQueryParamFilter("action", "CreateDatasetVersion").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateJob
//
// PARAMS:
//   - request: the arguments to CreateJob
//
// RETURNS:
//   - CreateJobResponse: The return type of the CreateJob interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateJob(request *CreateJobRequest) (*CreateJobResponse, error) {
	result := &CreateJobResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateJobUri()).
		WithQueryParamFilter("action", "CreateJob").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateModel
//
// PARAMS:
//   - request: the arguments to CreateModel
//
// RETURNS:
//   - CreateModelResponse: The return type of the CreateModel interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateModel(request *CreateModelRequest) (*CreateModelResponse, error) {
	result := &CreateModelResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateModelUri()).
		WithQueryParamFilter("action", "CreateModel").
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateModelVersion
//
// PARAMS:
//   - request: the arguments to CreateModelVersion
//
// RETURNS:
//   - CreateModelVersionResponse: The return type of the CreateModelVersion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateModelVersion(request *CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	result := &CreateModelVersionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateModelVersionUri()).
		WithQueryParamFilter("action", "CreateModelVersion").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteDataset
//
// PARAMS:
//   - request: the arguments to DeleteDataset
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDataset(request *DeleteDatasetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteDatasetUri()).
		WithQueryParamFilter("action", "DeleteDataset").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithHeader("Version", "v2").
		Do()
}

// DeleteDatasetVersion
//
// PARAMS:
//   - request: the arguments to DeleteDatasetVersion
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDatasetVersion(request *DeleteDatasetVersionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteDatasetVersionUri()).
		WithQueryParamFilter("action", "DeleteDatasetVersion").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithHeader("Version", "v2").
		Do()
}

// DeleteJob
//
// PARAMS:
//   - request: the arguments to DeleteJob
//
// RETURNS:
//   - DeleteJobResponse: The return type of the DeleteJob interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteJob(request *DeleteJobRequest) (*DeleteJobResponse, error) {
	result := &DeleteJobResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteJobUri()).
		WithQueryParamFilter("action", "DeleteJob").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteModel
//
// PARAMS:
//   - request: the arguments to DeleteModel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteModel(request *DeleteModelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteModelUri()).
		WithQueryParamFilter("action", "DeleteModel").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithHeader("Version", "v2").
		Do()
}

// DeleteModelVersion
//
// PARAMS:
//   - request: the arguments to DeleteModelVersion
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteModelVersion(request *DeleteModelVersionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteModelVersionUri()).
		WithQueryParamFilter("action", "DeleteModelVersion").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithHeader("Version", "v2").
		Do()
}

// DescribeDataset
//
// PARAMS:
//   - request: the arguments to DescribeDataset
//
// RETURNS:
//   - DescribeDatasetResponse: The return type of the DescribeDataset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDataset(request *DescribeDatasetRequest) (*DescribeDatasetResponse, error) {
	result := &DescribeDatasetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeDatasetUri()).
		WithQueryParamFilter("action", "DescribeDataset").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDatasetVersion
//
// PARAMS:
//   - request: the arguments to DescribeDatasetVersion
//
// RETURNS:
//   - DescribeDatasetVersionResponse: The return type of the DescribeDatasetVersion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDatasetVersion(request *DescribeDatasetVersionRequest) (*DescribeDatasetVersionResponse, error) {
	result := &DescribeDatasetVersionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeDatasetVersionUri()).
		WithQueryParamFilter("action", "DescribeDatasetVersion").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDatasetVersions
//
// PARAMS:
//   - request: the arguments to DescribeDatasetVersions
//
// RETURNS:
//   - DescribeDatasetVersionsResponse: The return type of the DescribeDatasetVersions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDatasetVersions(request *DescribeDatasetVersionsRequest) (*DescribeDatasetVersionsResponse, error) {
	result := &DescribeDatasetVersionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeDatasetVersionsUri()).
		WithQueryParamFilter("action", "DescribeDatasetVersions").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDatasets
//
// PARAMS:
//   - request: the arguments to DescribeDatasets
//
// RETURNS:
//   - DescribeDatasetsResponse: The return type of the DescribeDatasets interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDatasets(request *DescribeDatasetsRequest) (*DescribeDatasetsResponse, error) {
	result := &DescribeDatasetsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeDatasetsUri()).
		WithQueryParamFilter("action", "DescribeDatasets").
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("storageType", util.StringValue(request.StorageType)).
		WithQueryParamFilter("storageInstances", util.StringValue(request.StorageInstances)).
		WithQueryParamFilter("importFormat", util.StringValue(request.ImportFormat)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJob
//
// PARAMS:
//   - request: the arguments to DescribeJob
//
// RETURNS:
//   - DescribeJobResponse: The return type of the DescribeJob interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJob(request *DescribeJobRequest) (*DescribeJobResponse, error) {
	result := &DescribeJobResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobUri()).
		WithQueryParamFilter("action", "DescribeJob").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJobEvents
//
// PARAMS:
//   - request: the arguments to DescribeJobEvents
//
// RETURNS:
//   - DescribeJobEventsResponse: The return type of the DescribeJobEvents interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJobEvents(request *DescribeJobEventsRequest) (*DescribeJobEventsResponse, error) {
	result := &DescribeJobEventsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobEventsUri()).
		WithQueryParamFilter("action", "DescribeJobEvents").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJobLogs
//
// PARAMS:
//   - request: the arguments to DescribeJobLogs
//
// RETURNS:
//   - DescribeJobLogsResponse: The return type of the DescribeJobLogs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJobLogs(request *DescribeJobLogsRequest) (*DescribeJobLogsResponse, error) {
	result := &DescribeJobLogsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobLogsUri()).
		WithQueryParamFilter("action", "DescribeJobLogs").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJobMetrics
//
// PARAMS:
//   - request: the arguments to DescribeJobMetrics
//
// RETURNS:
//   - DescribeJobMetricsResponse: The return type of the DescribeJobMetrics interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJobMetrics(request *DescribeJobMetricsRequest) (*DescribeJobMetricsResponse, error) {
	result := &DescribeJobMetricsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobMetricsUri()).
		WithQueryParamFilter("action", "DescribeJobMetrics").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJobNodes
//
// PARAMS:
//   - request: the arguments to DescribeJobNodes
//
// RETURNS:
//   - DescribeJobNodesResponse: The return type of the DescribeJobNodes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJobNodes(request *DescribeJobNodesRequest) (*DescribeJobNodesResponse, error) {
	result := &DescribeJobNodesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobNodesUri()).
		WithQueryParamFilter("action", "DescribeJobNodes").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJobWebterminal
//
// PARAMS:
//   - request: the arguments to DescribeJobWebterminal
//
// RETURNS:
//   - DescribeJobWebterminalResponse: The return type of the DescribeJobWebterminal interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJobWebterminal(request *DescribeJobWebterminalRequest) (*DescribeJobWebterminalResponse, error) {
	result := &DescribeJobWebterminalResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobWebterminalUri()).
		WithQueryParamFilter("action", "DescribeJobWebterminal").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeJobs
//
// PARAMS:
//   - request: the arguments to DescribeJobs
//
// RETURNS:
//   - DescribeJobsResponse: The return type of the DescribeJobs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (*DescribeJobsResponse, error) {
	result := &DescribeJobsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeJobsUri()).
		WithQueryParamFilter("action", "DescribeJobs").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeModel
//
// PARAMS:
//   - request: the arguments to DescribeModel
//
// RETURNS:
//   - DescribeModelResponse: The return type of the DescribeModel interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeModel(request *DescribeModelRequest) (*DescribeModelResponse, error) {
	result := &DescribeModelResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeModelUri()).
		WithQueryParamFilter("action", "DescribeModel").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeModelVersion
//
// PARAMS:
//   - request: the arguments to DescribeModelVersion
//
// RETURNS:
//   - DescribeModelVersionResponse: The return type of the DescribeModelVersion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeModelVersion(request *DescribeModelVersionRequest) (*DescribeModelVersionResponse, error) {
	result := &DescribeModelVersionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeModelVersionUri()).
		WithQueryParamFilter("action", "DescribeModelVersion").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeModelVersions
//
// PARAMS:
//   - request: the arguments to DescribeModelVersions
//
// RETURNS:
//   - DescribeModelVersionsResponse: The return type of the DescribeModelVersions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeModelVersions(request *DescribeModelVersionsRequest) (*DescribeModelVersionsResponse, error) {
	result := &DescribeModelVersionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeModelVersionsUri()).
		WithQueryParamFilter("action", "DescribeModelVersions").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeModels
//
// PARAMS:
//   - request: the arguments to DescribeModels
//
// RETURNS:
//   - DescribeModelsResponse: The return type of the DescribeModels interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeModels(request *DescribeModelsRequest) (*DescribeModelsResponse, error) {
	result := &DescribeModelsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeModelsUri()).
		WithQueryParamFilter("action", "DescribeModels").
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithHeader("Version", "v2").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribePodEvents
//
// PARAMS:
//   - request: the arguments to DescribePodEvents
//
// RETURNS:
//   - DescribePodEventsResponse: The return type of the DescribePodEvents interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribePodEvents(request *DescribePodEventsRequest) (*DescribePodEventsResponse, error) {
	result := &DescribePodEventsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribePodEventsUri()).
		WithQueryParamFilter("action", "DescribePodEvents").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyDataset
//
// PARAMS:
//   - request: the arguments to ModifyDataset
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyDataset(request *ModifyDatasetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyDatasetUri()).
		WithQueryParamFilter("action", "ModifyDataset").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithHeader("Version", "v2").
		WithBody(request).
		Do()
}

// ModifyJob
//
// PARAMS:
//   - request: the arguments to ModifyJob
//
// RETURNS:
//   - ModifyJobResponse: The return type of the ModifyJob interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyJob(request *ModifyJobRequest) (*ModifyJobResponse, error) {
	result := &ModifyJobResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyJobUri()).
		WithQueryParamFilter("action", "ModifyJob").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyModel
//
// PARAMS:
//   - request: the arguments to ModifyModel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyModel(request *ModifyModelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyModelUri()).
		WithQueryParamFilter("action", "ModifyModel").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithHeader("Version", "v2").
		WithBody(request).
		Do()
}

// StopJob
//
// PARAMS:
//   - request: the arguments to StopJob
//
// RETURNS:
//   - StopJobResponse: The return type of the StopJob interface.
//   - error: nil if success otherwise the specific error
func (c *Client) StopJob(request *StopJobRequest) (*StopJobResponse, error) {
	result := &StopJobResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getStopJobUri()).
		WithQueryParamFilter("action", "StopJob").
		WithQueryParamFilter("resourcePoolId", "xxxx").
		WithQueryParamFilter("queueID", "xxxx").
		WithQueryParamFilter("resourcePoolId", util.StringValue(request.ResourcePoolId)).
		WithQueryParamFilter("queueID", util.StringValue(request.QueueID)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
