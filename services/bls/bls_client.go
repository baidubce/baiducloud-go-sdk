package bls

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"

	VERSION_V2 = "v2"
)

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
