package cloudassistant

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// ActionList
//
// PARAMS:
//   - request: the arguments to ActionList
//
// RETURNS:
//   - ActionListResponse: The return type of the ActionList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ActionList(request *ActionListRequest) (*ActionListResponse, error) {
	result := &ActionListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getActionListUri(VERSION_V1)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ActionLog
//
// PARAMS:
//   - request: the arguments to ActionLog
//
// RETURNS:
//   - ActionLogResponse: The return type of the ActionLog interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ActionLog(request *ActionLogRequest) (*ActionLogResponse, error) {
	result := &ActionLogResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getActionLogUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ActionRun
//
// PARAMS:
//   - request: the arguments to ActionRun
//
// RETURNS:
//   - ActionRunResponse: The return type of the ActionRun interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ActionRun(request *ActionRunRequest) (*ActionRunResponse, error) {
	result := &ActionRunResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getActionRunUri(VERSION_V1)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ActionRunList
//
// PARAMS:
//   - request: the arguments to ActionRunList
//
// RETURNS:
//   - ActionRunListResponse: The return type of the ActionRunList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ActionRunList(request *ActionRunListRequest) (*ActionRunListResponse, error) {
	result := &ActionRunListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getActionRunListUri(VERSION_V1)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchGetAgent
//
// PARAMS:
//   - request: the arguments to BatchGetAgent
//
// RETURNS:
//   - BatchGetAgentResponse: The return type of the BatchGetAgent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchGetAgent(request *BatchGetAgentRequest) (*BatchGetAgentResponse, error) {
	result := &BatchGetAgentResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchGetAgentUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAction
//
// PARAMS:
//   - request: the arguments to CreateAction
//
// RETURNS:
//   - CreateActionResponse: The return type of the CreateAction interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAction(request *CreateActionRequest) (*CreateActionResponse, error) {
	result := &CreateActionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateActionUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAction
//
// PARAMS:
//   - request: the arguments to DeleteAction
//
// RETURNS:
//   - DeleteActionResponse: The return type of the DeleteAction interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAction(request *DeleteActionRequest) (*DeleteActionResponse, error) {
	result := &DeleteActionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteActionUri(VERSION_V1, util.StringValue(request.Id))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAction
//
// PARAMS:
//   - request: the arguments to GetAction
//
// RETURNS:
//   - GetActionResponse: The return type of the GetAction interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAction(request *GetActionRequest) (*GetActionResponse, error) {
	result := &GetActionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetActionUri(VERSION_V1)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetActionRun
//
// PARAMS:
//   - request: the arguments to GetActionRun
//
// RETURNS:
//   - GetActionRunResponse: The return type of the GetActionRun interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetActionRun(request *GetActionRunRequest) (*GetActionRunResponse, error) {
	result := &GetActionRunResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetActionRunUri(VERSION_V1)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("withLog", util.StringValue(request.WithLog)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("childRunState", util.StringValue(request.ChildRunState)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAction
//
// PARAMS:
//   - request: the arguments to UpdateAction
//
// RETURNS:
//   - UpdateActionResponse: The return type of the UpdateAction interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAction(request *UpdateActionRequest) (*UpdateActionResponse, error) {
	result := &UpdateActionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateActionUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
