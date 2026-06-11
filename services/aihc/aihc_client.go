package aihc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// CreateADatasetV2
//
// PARAMS:
//   - request: the arguments to CreateADatasetV2
//
// RETURNS:
//   - CreateADatasetV2Response: The return type of the CreateADatasetV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateADatasetV2(request *CreateADatasetV2Request) (*CreateADatasetV2Response, error) {
	result := &CreateADatasetV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateADatasetV2Uri()).
		WithQueryParamFilter("action", "CreateDataset").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAModelV2
//
// PARAMS:
//   - request: the arguments to CreateAModelV2
//
// RETURNS:
//   - CreateAModelV2Response: The return type of the CreateAModelV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAModelV2(request *CreateAModelV2Request) (*CreateAModelV2Response, error) {
	result := &CreateAModelV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAModelV2Uri()).
		WithQueryParamFilter("action", "CreateModel").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDatasetVersionV2
//
// PARAMS:
//   - request: the arguments to CreateDatasetVersionV2
//
// RETURNS:
//   - CreateDatasetVersionV2Response: The return type of the CreateDatasetVersionV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDatasetVersionV2(request *CreateDatasetVersionV2Request) (*CreateDatasetVersionV2Response, error) {
	result := &CreateDatasetVersionV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDatasetVersionV2Uri()).
		WithQueryParamFilter("action", "CreateDatasetVersion").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteDatasetV2
//
// PARAMS:
//   - request: the arguments to DeleteDatasetV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDatasetV2(request *DeleteDatasetV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteDatasetV2Uri()).
		WithQueryParamFilter("action", "DeleteDataset").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		Do()
}

// DeleteDatasetVersionV2
//
// PARAMS:
//   - request: the arguments to DeleteDatasetVersionV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDatasetVersionV2(request *DeleteDatasetVersionV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteDatasetVersionV2Uri()).
		WithQueryParamFilter("action", "DeleteDatasetVersion").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		Do()
}

// DeleteModelV2
//
// PARAMS:
//   - request: the arguments to DeleteModelV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteModelV2(request *DeleteModelV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteModelV2Uri()).
		WithQueryParamFilter("action", "DeleteModel").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		Do()
}

// DeleteModelVersionV2
//
// PARAMS:
//   - request: the arguments to DeleteModelVersionV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteModelVersionV2(request *DeleteModelVersionV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteModelVersionV2Uri()).
		WithQueryParamFilter("action", "DeleteModelVersion").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		Do()
}

// GetAListOfModelVersionsV2
//
// PARAMS:
//   - request: the arguments to GetAListOfModelVersionsV2
//
// RETURNS:
//   - GetAListOfModelVersionsV2Response: The return type of the GetAListOfModelVersionsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAListOfModelVersionsV2(request *GetAListOfModelVersionsV2Request) (*GetAListOfModelVersionsV2Response, error) {
	result := &GetAListOfModelVersionsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAListOfModelVersionsV2Uri()).
		WithQueryParamFilter("action", "DescribeModelVersions").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDatasetDetailsV2
//
// PARAMS:
//   - request: the arguments to GetDatasetDetailsV2
//
// RETURNS:
//   - GetDatasetDetailsV2Response: The return type of the GetDatasetDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDatasetDetailsV2(request *GetDatasetDetailsV2Request) (*GetDatasetDetailsV2Response, error) {
	result := &GetDatasetDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDatasetDetailsV2Uri()).
		WithQueryParamFilter("action", "DescribeDataset").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDatasetVersionDetailsV2
//
// PARAMS:
//   - request: the arguments to GetDatasetVersionDetailsV2
//
// RETURNS:
//   - GetDatasetVersionDetailsV2Response: The return type of the GetDatasetVersionDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDatasetVersionDetailsV2(request *GetDatasetVersionDetailsV2Request) (*GetDatasetVersionDetailsV2Response, error) {
	result := &GetDatasetVersionDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDatasetVersionDetailsV2Uri()).
		WithQueryParamFilter("action", "DescribeDatasetVersion").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetModelDetailsV2
//
// PARAMS:
//   - request: the arguments to GetModelDetailsV2
//
// RETURNS:
//   - GetModelDetailsV2Response: The return type of the GetModelDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetModelDetailsV2(request *GetModelDetailsV2Request) (*GetModelDetailsV2Response, error) {
	result := &GetModelDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetModelDetailsV2Uri()).
		WithQueryParamFilter("action", "DescribeModel").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetModelListV2
//
// PARAMS:
//   - request: the arguments to GetModelListV2
//
// RETURNS:
//   - GetModelListV2Response: The return type of the GetModelListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetModelListV2(request *GetModelListV2Request) (*GetModelListV2Response, error) {
	result := &GetModelListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetModelListV2Uri()).
		WithQueryParamFilter("action", "DescribeModels").
		WithQueryParamFilter("keyword", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetModelVersionDetailsV2
//
// PARAMS:
//   - request: the arguments to GetModelVersionDetailsV2
//
// RETURNS:
//   - GetModelVersionDetailsV2Response: The return type of the GetModelVersionDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetModelVersionDetailsV2(request *GetModelVersionDetailsV2Request) (*GetModelVersionDetailsV2Response, error) {
	result := &GetModelVersionDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetModelVersionDetailsV2Uri()).
		WithQueryParamFilter("action", "DescribeModelVersion").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("versionId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithQueryParamFilter("versionId", util.StringValue(request.VersionId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyDatasetV2
//
// PARAMS:
//   - request: the arguments to ModifyDatasetV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyDatasetV2(request *ModifyDatasetV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyDatasetV2Uri()).
		WithQueryParamFilter("action", "ModifyDataset").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithBody(request).
		Do()
}

// ModifyTheModelV2
//
// PARAMS:
//   - request: the arguments to ModifyTheModelV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTheModelV2(request *ModifyTheModelV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyTheModelV2Uri()).
		WithQueryParamFilter("action", "ModifyModel").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithBody(request).
		Do()
}

// NewModelVersionV2
//
// PARAMS:
//   - request: the arguments to NewModelVersionV2
//
// RETURNS:
//   - NewModelVersionV2Response: The return type of the NewModelVersionV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) NewModelVersionV2(request *NewModelVersionV2Request) (*NewModelVersionV2Response, error) {
	result := &NewModelVersionV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNewModelVersionV2Uri()).
		WithQueryParamFilter("action", "CreateModelVersion").
		WithQueryParamFilter("modelId", "xxx").
		WithQueryParamFilter("modelId", util.StringValue(request.ModelId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RetrieveTheDatasetListV2
//
// PARAMS:
//   - request: the arguments to RetrieveTheDatasetListV2
//
// RETURNS:
//   - RetrieveTheDatasetListV2Response: The return type of the RetrieveTheDatasetListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveTheDatasetListV2(request *RetrieveTheDatasetListV2Request) (*RetrieveTheDatasetListV2Response, error) {
	result := &RetrieveTheDatasetListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveTheDatasetListV2Uri()).
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
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RetrieveTheDatasetVersionListV2
//
// PARAMS:
//   - request: the arguments to RetrieveTheDatasetVersionListV2
//
// RETURNS:
//   - RetrieveTheDatasetVersionListV2Response: The return type of the RetrieveTheDatasetVersionListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveTheDatasetVersionListV2(request *RetrieveTheDatasetVersionListV2Request) (*RetrieveTheDatasetVersionListV2Response, error) {
	result := &RetrieveTheDatasetVersionListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveTheDatasetVersionListV2Uri()).
		WithQueryParamFilter("action", "DescribeDatasetVersions").
		WithQueryParamFilter("datasetId", "xxx").
		WithQueryParamFilter("pageNumber", "xxx").
		WithQueryParamFilter("pageSize", "xxx").
		WithQueryParamFilter("datasetId", util.StringValue(request.DatasetId)).
		WithQueryParamFilter("pageNumber", util.Int32Value(request.PageNumber)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
