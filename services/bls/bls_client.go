package bls

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"

	VERSION_V2 = "v2"

	VERSION_V3 = "v3"
)

// AsyncSearch
//
// PARAMS:
//   - request: the arguments to AsyncSearch
//
// RETURNS:
//   - AsyncSearchResponse: The return type of the AsyncSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AsyncSearch(request *AsyncSearchRequest) (*AsyncSearchResponse, error) {
	result := &AsyncSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAsyncSearchUri(util.StringValue(request.Name))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchGetLogStore
//
// PARAMS:
//   - request: the arguments to BatchGetLogStore
//
// RETURNS:
//   - BatchGetLogStoreResponse: The return type of the BatchGetLogStore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchGetLogStore(request *BatchGetLogStoreRequest) (*BatchGetLogStoreResponse, error) {
	result := &BatchGetLogStoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchGetLogStoreUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BulkDeleteLogShipper
//
// PARAMS:
//   - request: the arguments to BulkDeleteLogShipper
//
// RETURNS:
//   - BulkDeleteLogShipperResponse: The return type of the BulkDeleteLogShipper interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BulkDeleteLogShipper(request *BulkDeleteLogShipperRequest) (*BulkDeleteLogShipperResponse, error) {
	result := &BulkDeleteLogShipperResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getBulkDeleteLogShipperUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BulkSetLogShipperStatus
//
// PARAMS:
//   - request: the arguments to BulkSetLogShipperStatus
//
// RETURNS:
//   - BulkSetLogShipperStatusResponse: The return type of the BulkSetLogShipperStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BulkSetLogShipperStatus(request *BulkSetLogShipperStatusRequest) (*BulkSetLogShipperStatusResponse, error) {
	result := &BulkSetLogShipperStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBulkSetLogShipperStatusUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAlarmPolicy
//
// PARAMS:
//   - request: the arguments to CreateAlarmPolicy
//
// RETURNS:
//   - CreateAlarmPolicyResponse: The return type of the CreateAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (*CreateAlarmPolicyResponse, error) {
	result := &CreateAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAlarmPolicyUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDownloadTask
//
// PARAMS:
//   - request: the arguments to CreateDownloadTask
//
// RETURNS:
//   - CreateDownloadTaskResponse: The return type of the CreateDownloadTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDownloadTask(request *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error) {
	result := &CreateDownloadTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDownloadTaskUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFastQuery
//
// PARAMS:
//   - request: the arguments to CreateFastQuery
//
// RETURNS:
//   - CreateFastQueryResponse: The return type of the CreateFastQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFastQuery(request *CreateFastQueryRequest) (*CreateFastQueryResponse, error) {
	result := &CreateFastQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateFastQueryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIndex
//
// PARAMS:
//   - request: the arguments to CreateIndex
//
// RETURNS:
//   - CreateIndexResponse: The return type of the CreateIndex interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error) {
	result := &CreateIndexResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIndexUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLogShipper
//
// PARAMS:
//   - request: the arguments to CreateLogShipper
//
// RETURNS:
//   - CreateLogShipperResponse: The return type of the CreateLogShipper interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateLogShipper(request *CreateLogShipperRequest) (*CreateLogShipperResponse, error) {
	result := &CreateLogShipperResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateLogShipperUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLogStore
//
// PARAMS:
//   - request: the arguments to CreateLogStore
//
// RETURNS:
//   - CreateLogStoreResponse: The return type of the CreateLogStore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateLogStore(request *CreateLogStoreRequest) (*CreateLogStoreResponse, error) {
	result := &CreateLogStoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateLogStoreUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLogStoreTemplate
//
// PARAMS:
//   - request: the arguments to CreateLogStoreTemplate
//
// RETURNS:
//   - CreateLogStoreTemplateResponse: The return type of the CreateLogStoreTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateLogStoreTemplate(request *CreateLogStoreTemplateRequest) (*CreateLogStoreTemplateResponse, error) {
	result := &CreateLogStoreTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateLogStoreTemplateUri(VERSION_V3)).
		WithQueryParamFilter("action", "CreateLogStoreTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLogStoreView
//
// PARAMS:
//   - request: the arguments to CreateLogStoreView
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateLogStoreView(request *CreateLogStoreViewRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateLogStoreViewUri(VERSION_V3)).
		WithQueryParamFilter("action", "CreateLogStoreView").
		WithBody(request).
		Do()
}

// CreateProject
//
// PARAMS:
//   - request: the arguments to CreateProject
//
// RETURNS:
//   - CreateProjectResponse: The return type of the CreateProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateProject(request *CreateProjectRequest) (*CreateProjectResponse, error) {
	result := &CreateProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateProjectUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTask
//
// PARAMS:
//   - request: the arguments to CreateTask
//
// RETURNS:
//   - CreateTaskResponse: The return type of the CreateTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateTask(request *CreateTaskRequest) (*CreateTaskResponse, error) {
	result := &CreateTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTaskUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAlarmPolicy
//
// PARAMS:
//   - request: the arguments to DeleteAlarmPolicy
//
// RETURNS:
//   - DeleteAlarmPolicyResponse: The return type of the DeleteAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (*DeleteAlarmPolicyResponse, error) {
	result := &DeleteAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAlarmPolicyUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteDownloadTask
//
// PARAMS:
//   - request: the arguments to DeleteDownloadTask
//
// RETURNS:
//   - DeleteDownloadTaskResponse: The return type of the DeleteDownloadTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteDownloadTask(request *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error) {
	result := &DeleteDownloadTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDownloadTaskUri(VERSION_V2, util.StringValue(request.Uuid))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteFastQuery
//
// PARAMS:
//   - request: the arguments to DeleteFastQuery
//
// RETURNS:
//   - DeleteFastQueryResponse: The return type of the DeleteFastQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteFastQuery(request *DeleteFastQueryRequest) (*DeleteFastQueryResponse, error) {
	result := &DeleteFastQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteFastQueryUri(VERSION_V1, util.StringValue(request.FastQueryName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteIndex
//
// PARAMS:
//   - request: the arguments to DeleteIndex
//
// RETURNS:
//   - DeleteIndexResponse: The return type of the DeleteIndex interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	result := &DeleteIndexResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIndexUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteLogStore
//
// PARAMS:
//   - request: the arguments to DeleteLogStore
//
// RETURNS:
//   - DeleteLogStoreResponse: The return type of the DeleteLogStore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteLogStore(request *DeleteLogStoreRequest) (*DeleteLogStoreResponse, error) {
	result := &DeleteLogStoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteLogStoreUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteLogStoreTemplates
//
// PARAMS:
//   - request: the arguments to DeleteLogStoreTemplates
//
// RETURNS:
//   - DeleteLogStoreTemplatesResponse: The return type of the DeleteLogStoreTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteLogStoreTemplates(request *DeleteLogStoreTemplatesRequest) (*DeleteLogStoreTemplatesResponse, error) {
	result := &DeleteLogStoreTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteLogStoreTemplatesUri(VERSION_V3)).
		WithQueryParamFilter("action", "DeleteLogStoreTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteLogStoreView
//
// PARAMS:
//   - request: the arguments to DeleteLogStoreView
//
// RETURNS:
//   - DeleteLogStoreViewResponse: The return type of the DeleteLogStoreView interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteLogStoreView(request *DeleteLogStoreViewRequest) (*DeleteLogStoreViewResponse, error) {
	result := &DeleteLogStoreViewResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteLogStoreViewUri(VERSION_V3)).
		WithQueryParamFilter("action", "DeleteLogStoreView").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteProject
//
// PARAMS:
//   - request: the arguments to DeleteProject
//
// RETURNS:
//   - DeleteProjectResponse: The return type of the DeleteProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteProject(request *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	result := &DeleteProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteProjectUri(VERSION_V1, util.StringValue(request.Uuid))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSingleLogShipper
//
// PARAMS:
//   - request: the arguments to DeleteSingleLogShipper
//
// RETURNS:
//   - DeleteSingleLogShipperResponse: The return type of the DeleteSingleLogShipper interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSingleLogShipper(request *DeleteSingleLogShipperRequest) (*DeleteSingleLogShipperResponse, error) {
	result := &DeleteSingleLogShipperResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSingleLogShipperUri(VERSION_V1, util.StringValue(request.LogShipperID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmPolicy
//
// PARAMS:
//   - request: the arguments to DescribeAlarmPolicy
//
// RETURNS:
//   - DescribeAlarmPolicyResponse: The return type of the DescribeAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) (*DescribeAlarmPolicyResponse, error) {
	result := &DescribeAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAlarmPolicyUri(VERSION_V1)).
		WithQueryParamFilter("PolicyName", util.StringValue(request.PolicyName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmRecord
//
// PARAMS:
//   - request: the arguments to DescribeAlarmRecord
//
// RETURNS:
//   - DescribeAlarmRecordResponse: The return type of the DescribeAlarmRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmRecord(request *DescribeAlarmRecordRequest) (*DescribeAlarmRecordResponse, error) {
	result := &DescribeAlarmRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAlarmRecordUri(VERSION_V1)).
		WithQueryParamFilter("alarmId", util.StringValue(request.AlarmId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDownloadTask
//
// PARAMS:
//   - request: the arguments to DescribeDownloadTask
//
// RETURNS:
//   - DescribeDownloadTaskResponse: The return type of the DescribeDownloadTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDownloadTask(request *DescribeDownloadTaskRequest) (*DescribeDownloadTaskResponse, error) {
	result := &DescribeDownloadTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeDownloadTaskUri(VERSION_V2, util.StringValue(request.Uuid))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeFastQuery
//
// PARAMS:
//   - request: the arguments to DescribeFastQuery
//
// RETURNS:
//   - DescribeFastQueryResponse: The return type of the DescribeFastQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeFastQuery(request *DescribeFastQueryRequest) (*DescribeFastQueryResponse, error) {
	result := &DescribeFastQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeFastQueryUri(VERSION_V1, util.StringValue(request.FastQueryName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeIndex
//
// PARAMS:
//   - request: the arguments to DescribeIndex
//
// RETURNS:
//   - DescribeIndexResponse: The return type of the DescribeIndex interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeIndex(request *DescribeIndexRequest) (*DescribeIndexResponse, error) {
	result := &DescribeIndexResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeIndexUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLogStore
//
// PARAMS:
//   - request: the arguments to DescribeLogStore
//
// RETURNS:
//   - DescribeLogStoreResponse: The return type of the DescribeLogStore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLogStore(request *DescribeLogStoreRequest) (*DescribeLogStoreResponse, error) {
	result := &DescribeLogStoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeLogStoreUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLogStoreTemplate
//
// PARAMS:
//   - request: the arguments to DescribeLogStoreTemplate
//
// RETURNS:
//   - DescribeLogStoreTemplateResponse: The return type of the DescribeLogStoreTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLogStoreTemplate(request *DescribeLogStoreTemplateRequest) (*DescribeLogStoreTemplateResponse, error) {
	result := &DescribeLogStoreTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLogStoreTemplateUri(VERSION_V3)).
		WithQueryParamFilter("action", "DescribeLogStoreTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLogStoreTemplates
//
// PARAMS:
//   - request: the arguments to DescribeLogStoreTemplates
//
// RETURNS:
//   - DescribeLogStoreTemplatesResponse: The return type of the DescribeLogStoreTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLogStoreTemplates(request *DescribeLogStoreTemplatesRequest) (*DescribeLogStoreTemplatesResponse, error) {
	result := &DescribeLogStoreTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLogStoreTemplatesUri(VERSION_V3)).
		WithQueryParamFilter("action", "DescribeLogStoreTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLogStoreView
//
// PARAMS:
//   - request: the arguments to DescribeLogStoreView
//
// RETURNS:
//   - DescribeLogStoreViewResponse: The return type of the DescribeLogStoreView interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLogStoreView(request *DescribeLogStoreViewRequest) (*DescribeLogStoreViewResponse, error) {
	result := &DescribeLogStoreViewResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLogStoreViewUri(VERSION_V3)).
		WithQueryParamFilter("action", "DescribeLogStoreView").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeProject
//
// PARAMS:
//   - request: the arguments to DescribeProject
//
// RETURNS:
//   - DescribeProjectResponse: The return type of the DescribeProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeProject(request *DescribeProjectRequest) (*DescribeProjectResponse, error) {
	result := &DescribeProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeProjectUri(VERSION_V1, util.StringValue(request.Uuid))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DisableAlarmPolicy
//
// PARAMS:
//   - request: the arguments to DisableAlarmPolicy
//
// RETURNS:
//   - DisableAlarmPolicyResponse: The return type of the DisableAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DisableAlarmPolicy(request *DisableAlarmPolicyRequest) (*DisableAlarmPolicyResponse, error) {
	result := &DisableAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDisableAlarmPolicyUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EnableAlarmPolicy
//
// PARAMS:
//   - request: the arguments to EnableAlarmPolicy
//
// RETURNS:
//   - EnableAlarmPolicyResponse: The return type of the EnableAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) EnableAlarmPolicy(request *EnableAlarmPolicyRequest) (*EnableAlarmPolicyResponse, error) {
	result := &EnableAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEnableAlarmPolicyUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FieldCaps
//
// PARAMS:
//   - request: the arguments to FieldCaps
//
// RETURNS:
//   - FieldCapsResponse: The return type of the FieldCaps interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FieldCaps(request *FieldCapsRequest) (*FieldCapsResponse, error) {
	result := &FieldCapsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFieldCapsUri(util.StringValue(request.Name))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDownloadTaskLink
//
// PARAMS:
//   - request: the arguments to GetDownloadTaskLink
//
// RETURNS:
//   - GetDownloadTaskLinkResponse: The return type of the GetDownloadTaskLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDownloadTaskLink(request *GetDownloadTaskLinkRequest) (*GetDownloadTaskLinkResponse, error) {
	result := &GetDownloadTaskLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDownloadTaskLinkUri(VERSION_V2, util.StringValue(request.Uuid))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLogShipper
//
// PARAMS:
//   - request: the arguments to GetLogShipper
//
// RETURNS:
//   - GetLogShipperResponse: The return type of the GetLogShipper interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetLogShipper(request *GetLogShipperRequest) (*GetLogShipperResponse, error) {
	result := &GetLogShipperResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetLogShipperUri(VERSION_V1, util.StringValue(request.LogShipperID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlarmExecutionStats
//
// PARAMS:
//   - request: the arguments to ListAlarmExecutionStats
//
// RETURNS:
//   - ListAlarmExecutionStatsResponse: The return type of the ListAlarmExecutionStats interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlarmExecutionStats(request *ListAlarmExecutionStatsRequest) (*ListAlarmExecutionStatsResponse, error) {
	result := &ListAlarmExecutionStatsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListAlarmExecutionStatsUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlarmExecutions
//
// PARAMS:
//   - request: the arguments to ListAlarmExecutions
//
// RETURNS:
//   - ListAlarmExecutionsResponse: The return type of the ListAlarmExecutions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlarmExecutions(request *ListAlarmExecutionsRequest) (*ListAlarmExecutionsResponse, error) {
	result := &ListAlarmExecutionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListAlarmExecutionsUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlarmPolicy
//
// PARAMS:
//   - request: the arguments to ListAlarmPolicy
//
// RETURNS:
//   - ListAlarmPolicyResponse: The return type of the ListAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlarmPolicy(request *ListAlarmPolicyRequest) (*ListAlarmPolicyResponse, error) {
	result := &ListAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListAlarmPolicyUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlarmRecord
//
// PARAMS:
//   - request: the arguments to ListAlarmRecord
//
// RETURNS:
//   - ListAlarmRecordResponse: The return type of the ListAlarmRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlarmRecord(request *ListAlarmRecordRequest) (*ListAlarmRecordResponse, error) {
	result := &ListAlarmRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListAlarmRecordUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListDownloadTask
//
// PARAMS:
//   - request: the arguments to ListDownloadTask
//
// RETURNS:
//   - ListDownloadTaskResponse: The return type of the ListDownloadTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListDownloadTask(request *ListDownloadTaskRequest) (*ListDownloadTaskResponse, error) {
	result := &ListDownloadTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListDownloadTaskUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListFastQuery
//
// PARAMS:
//   - request: the arguments to ListFastQuery
//
// RETURNS:
//   - ListFastQueryResponse: The return type of the ListFastQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListFastQuery(request *ListFastQueryRequest) (*ListFastQueryResponse, error) {
	result := &ListFastQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListFastQueryUri(VERSION_V1)).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithQueryParamFilter("logStoreName", util.StringValue(request.LogStoreName)).
		WithQueryParamFilter("namePattern", util.StringValue(request.NamePattern)).
		WithQueryParamFilter("logStoreType", util.StringValue(request.LogStoreType)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogShipper
//
// PARAMS:
//   - request: the arguments to ListLogShipper
//
// RETURNS:
//   - ListLogShipperResponse: The return type of the ListLogShipper interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogShipper(request *ListLogShipperRequest) (*ListLogShipperResponse, error) {
	result := &ListLogShipperResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListLogShipperUri(VERSION_V1)).
		WithQueryParam("", "").
		WithQueryParamFilter("logShipperID", util.StringValue(request.LogShipperID)).
		WithQueryParamFilter("logShipperName", util.StringValue(request.LogShipperName)).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithQueryParamFilter("logStoreName", util.StringValue(request.LogStoreName)).
		WithQueryParamFilter("destType", util.StringValue(request.DestType)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogShipperRecord
//
// PARAMS:
//   - request: the arguments to ListLogShipperRecord
//
// RETURNS:
//   - ListLogShipperRecordResponse: The return type of the ListLogShipperRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogShipperRecord(request *ListLogShipperRecordRequest) (*ListLogShipperRecordResponse, error) {
	result := &ListLogShipperRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListLogShipperRecordUri(VERSION_V1, util.StringValue(request.LogShipperID))).
		WithQueryParamFilter("sinceHours", util.Int32Value(request.SinceHours)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogStore
//
// PARAMS:
//   - request: the arguments to ListLogStore
//
// RETURNS:
//   - ListLogStoreResponse: The return type of the ListLogStore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogStore(request *ListLogStoreRequest) (*ListLogStoreResponse, error) {
	result := &ListLogStoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListLogStoreUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogStoreView
//
// PARAMS:
//   - request: the arguments to ListLogStoreView
//
// RETURNS:
//   - ListLogStoreViewResponse: The return type of the ListLogStoreView interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogStoreView(request *ListLogStoreViewRequest) (*ListLogStoreViewResponse, error) {
	result := &ListLogStoreViewResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListLogStoreViewUri(VERSION_V3)).
		WithQueryParamFilter("action", "DescribeLogStoreViews").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogStream
//
// PARAMS:
//   - request: the arguments to ListLogStream
//
// RETURNS:
//   - ListLogStreamResponse: The return type of the ListLogStream interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogStream(request *ListLogStreamRequest) (*ListLogStreamResponse, error) {
	result := &ListLogStreamResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListLogStreamUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithQueryParamFilter("namePattern", util.StringValue(request.NamePattern)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListProject
//
// PARAMS:
//   - request: the arguments to ListProject
//
// RETURNS:
//   - ListProjectResponse: The return type of the ListProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListProject(request *ListProjectRequest) (*ListProjectResponse, error) {
	result := &ListProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListProjectUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PullLogRecord
//
// PARAMS:
//   - request: the arguments to PullLogRecord
//
// RETURNS:
//   - PullLogRecordResponse: The return type of the PullLogRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PullLogRecord(request *PullLogRecordRequest) (*PullLogRecordResponse, error) {
	result := &PullLogRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPullLogRecordUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithQueryParamFilter("logStreamName", util.StringValue(request.LogStreamName)).
		WithQueryParamFilter("startDateTime", util.StringValue(request.StartDateTime)).
		WithQueryParamFilter("endDateTime", util.StringValue(request.EndDateTime)).
		WithQueryParamFilter("limit", util.Int32Value(request.Limit)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PushLogRecord
//
// PARAMS:
//   - request: the arguments to PushLogRecord
//
// RETURNS:
//   - PushLogRecordResponse: The return type of the PushLogRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PushLogRecord(request *PushLogRecordRequest) (*PushLogRecordResponse, error) {
	result := &PushLogRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPushLogRecordUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryLogHistogram
//
// PARAMS:
//   - request: the arguments to QueryLogHistogram
//
// RETURNS:
//   - QueryLogHistogramResponse: The return type of the QueryLogHistogram interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryLogHistogram(request *QueryLogHistogramRequest) (*QueryLogHistogramResponse, error) {
	result := &QueryLogHistogramResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryLogHistogramUri(VERSION_V2, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithQueryParamFilter("logStreamName", util.StringValue(request.LogStreamName)).
		WithQueryParamFilter("query", util.StringValue(request.Query)).
		WithQueryParamFilter("startDateTime", util.StringValue(request.StartDateTime)).
		WithQueryParamFilter("endDateTime", util.StringValue(request.EndDateTime)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryLogRecord
//
// PARAMS:
//   - request: the arguments to QueryLogRecord
//
// RETURNS:
//   - QueryLogRecordResponse: The return type of the QueryLogRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryLogRecord(request *QueryLogRecordRequest) (*QueryLogRecordResponse, error) {
	result := &QueryLogRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryLogRecordUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithQueryParamFilter("logStreamName", util.StringValue(request.LogStreamName)).
		WithQueryParamFilter("query", util.StringValue(request.Query)).
		WithQueryParamFilter("startDateTime", util.StringValue(request.StartDateTime)).
		WithQueryParamFilter("endDateTime", util.StringValue(request.EndDateTime)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("limit", util.Int32Value(request.Limit)).
		WithQueryParamFilter("sort", util.StringValue(request.Sort)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResolveIndex
//
// PARAMS:
//   - request: the arguments to ResolveIndex
//
// RETURNS:
//   - ResolveIndexResponse: The return type of the ResolveIndex interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResolveIndex(request *ResolveIndexRequest) (*ResolveIndexResponse, error) {
	result := &ResolveIndexResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getResolveIndexUri(util.StringValue(request.Name))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetSingleLogShipperStatus
//
// PARAMS:
//   - request: the arguments to SetSingleLogShipperStatus
//
// RETURNS:
//   - SetSingleLogShipperStatusResponse: The return type of the SetSingleLogShipperStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SetSingleLogShipperStatus(request *SetSingleLogShipperStatusRequest) (*SetSingleLogShipperStatusResponse, error) {
	result := &SetSingleLogShipperStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSetSingleLogShipperStatusUri(VERSION_V1, util.StringValue(request.LogShipperID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TermsEnum
//
// PARAMS:
//   - request: the arguments to TermsEnum
//
// RETURNS:
//   - TermsEnumResponse: The return type of the TermsEnum interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TermsEnum(request *TermsEnumRequest) (*TermsEnumResponse, error) {
	result := &TermsEnumResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTermsEnumUri(util.StringValue(request.Name))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAlarmPolicy
//
// PARAMS:
//   - request: the arguments to UpdateAlarmPolicy
//
// RETURNS:
//   - UpdateAlarmPolicyResponse: The return type of the UpdateAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlarmPolicy(request *UpdateAlarmPolicyRequest) (*UpdateAlarmPolicyResponse, error) {
	result := &UpdateAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAlarmPolicyUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFastQuery
//
// PARAMS:
//   - request: the arguments to UpdateFastQuery
//
// RETURNS:
//   - UpdateFastQueryResponse: The return type of the UpdateFastQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateFastQuery(request *UpdateFastQueryRequest) (*UpdateFastQueryResponse, error) {
	result := &UpdateFastQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateFastQueryUri(VERSION_V1, util.StringValue(request.Name))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIndex
//
// PARAMS:
//   - request: the arguments to UpdateIndex
//
// RETURNS:
//   - UpdateIndexResponse: The return type of the UpdateIndex interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateIndex(request *UpdateIndexRequest) (*UpdateIndexResponse, error) {
	result := &UpdateIndexResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIndexUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateLogShipper
//
// PARAMS:
//   - request: the arguments to UpdateLogShipper
//
// RETURNS:
//   - UpdateLogShipperResponse: The return type of the UpdateLogShipper interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateLogShipper(request *UpdateLogShipperRequest) (*UpdateLogShipperResponse, error) {
	result := &UpdateLogShipperResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateLogShipperUri(VERSION_V1, util.StringValue(request.LogShipperID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateLogStore
//
// PARAMS:
//   - request: the arguments to UpdateLogStore
//
// RETURNS:
//   - UpdateLogStoreResponse: The return type of the UpdateLogStore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateLogStore(request *UpdateLogStoreRequest) (*UpdateLogStoreResponse, error) {
	result := &UpdateLogStoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateLogStoreUri(VERSION_V1, util.StringValue(request.LogStoreName))).
		WithQueryParamFilter("project", util.StringValue(request.Project)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateLogStoreTemplate
//
// PARAMS:
//   - request: the arguments to UpdateLogStoreTemplate
//
// RETURNS:
//   - UpdateLogStoreTemplateResponse: The return type of the UpdateLogStoreTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateLogStoreTemplate(request *UpdateLogStoreTemplateRequest) (*UpdateLogStoreTemplateResponse, error) {
	result := &UpdateLogStoreTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateLogStoreTemplateUri(VERSION_V3)).
		WithQueryParamFilter("action", "UpdateLogStoreTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateLogStoreView
//
// PARAMS:
//   - request: the arguments to UpdateLogStoreView
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateLogStoreView(request *UpdateLogStoreViewRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateLogStoreViewUri(VERSION_V3)).
		WithQueryParamFilter("action", "UpdateLogStoreView").
		WithBody(request).
		Do()
}

// UpdateProject
//
// PARAMS:
//   - request: the arguments to UpdateProject
//
// RETURNS:
//   - UpdateProjectResponse: The return type of the UpdateProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateProject(request *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	result := &UpdateProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateProjectUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTask
//
// PARAMS:
//   - request: the arguments to UpdateTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTask(request *UpdateTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateTaskUri(util.StringValue(request.TaskId))).
		WithBody(request).
		Do()
}

// ValidateAlarmCondition
//
// PARAMS:
//   - request: the arguments to ValidateAlarmCondition
//
// RETURNS:
//   - ValidateAlarmConditionResponse: The return type of the ValidateAlarmCondition interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ValidateAlarmCondition(request *ValidateAlarmConditionRequest) (*ValidateAlarmConditionResponse, error) {
	result := &ValidateAlarmConditionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getValidateAlarmConditionUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ValidateAlarmPolicySQL
//
// PARAMS:
//   - request: the arguments to ValidateAlarmPolicySQL
//
// RETURNS:
//   - ValidateAlarmPolicySQLResponse: The return type of the ValidateAlarmPolicySQL interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ValidateAlarmPolicySQL(request *ValidateAlarmPolicySQLRequest) (*ValidateAlarmPolicySQLResponse, error) {
	result := &ValidateAlarmPolicySQLResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getValidateAlarmPolicySQLUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
