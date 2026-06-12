package aihc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

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
		WithQueryParamFilter("datasetId", "xxx").
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
		WithQueryParamFilter("modelId", "xxx").
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
		WithQueryParamFilter("datasetId", "xxx").
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
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithHeader("Version", "v2").
		Do()
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
		WithQueryParamFilter("modelId", "xxx").
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
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
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
		WithQueryParamFilter("datasetId", "xxx").
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
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
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
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
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
		WithQueryParamFilter("keyword", "xxx").
		WithQueryParamFilter("storageType", "xxx").
		WithQueryParamFilter("storageInstances", "xxx").
		WithQueryParamFilter("importFormat", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
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
		WithQueryParamFilter("modelId", "xxx").
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
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
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
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
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
		WithQueryParamFilter("keyword", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
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
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithHeader("Version", "v2").
		WithBody(request).
		Do()
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
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithHeader("Version", "v2").
		WithBody(request).
		Do()
}
