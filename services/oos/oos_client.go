package oos

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V2 = "v2"
)

// CheckTemplateV2
//
// PARAMS:
//   - request: the arguments to CheckTemplateV2
//
// RETURNS:
//   - CheckTemplateV2Response: The return type of the CheckTemplateV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CheckTemplateV2(request *CheckTemplateV2Request) (*CheckTemplateV2Response, error) {
	result := &CheckTemplateV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCheckTemplateV2Uri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateExecutionV2
//
// PARAMS:
//   - request: the arguments to CreateExecutionV2
//
// RETURNS:
//   - CreateExecutionV2Response: The return type of the CreateExecutionV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateExecutionV2(request *CreateExecutionV2Request) (*CreateExecutionV2Response, error) {
	result := &CreateExecutionV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateExecutionV2Uri(VERSION_V2)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTemplateV2
//
// PARAMS:
//   - request: the arguments to CreateTemplateV2
//
// RETURNS:
//   - CreateTemplateV2Response: The return type of the CreateTemplateV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateTemplateV2(request *CreateTemplateV2Request) (*CreateTemplateV2Response, error) {
	result := &CreateTemplateV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTemplateV2Uri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteTemplateV2
//
// PARAMS:
//   - request: the arguments to DeleteTemplateV2
//
// RETURNS:
//   - DeleteTemplateV2Response: The return type of the DeleteTemplateV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteTemplateV2(request *DeleteTemplateV2Request) (*DeleteTemplateV2Response, error) {
	result := &DeleteTemplateV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTemplateV2Uri(VERSION_V2)).
		WithQueryParamFilter("namespace", util.StringValue(request.Namespace)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetExecutionDetailV2
//
// PARAMS:
//   - request: the arguments to GetExecutionDetailV2
//
// RETURNS:
//   - GetExecutionDetailV2Response: The return type of the GetExecutionDetailV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetExecutionDetailV2(request *GetExecutionDetailV2Request) (*GetExecutionDetailV2Response, error) {
	result := &GetExecutionDetailV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetExecutionDetailV2Uri(VERSION_V2)).
		WithQueryParamFilter("namespace", util.StringValue(request.Namespace)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("withLog", util.StringValue(request.WithLog)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetExecutionListV2
//
// PARAMS:
//   - request: the arguments to GetExecutionListV2
//
// RETURNS:
//   - GetExecutionListV2Response: The return type of the GetExecutionListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetExecutionListV2(request *GetExecutionListV2Request) (*GetExecutionListV2Response, error) {
	result := &GetExecutionListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetExecutionListV2Uri(VERSION_V2)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOperatorListV2
//
// PARAMS:
//   - request: the arguments to GetOperatorListV2
//
// RETURNS:
//   - GetOperatorListV2Response: The return type of the GetOperatorListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetOperatorListV2(request *GetOperatorListV2Request) (*GetOperatorListV2Response, error) {
	result := &GetOperatorListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetOperatorListV2Uri(VERSION_V2)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskChildrenListV2
//
// PARAMS:
//   - request: the arguments to GetTaskChildrenListV2
//
// RETURNS:
//   - GetTaskChildrenListV2Response: The return type of the GetTaskChildrenListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTaskChildrenListV2(request *GetTaskChildrenListV2Request) (*GetTaskChildrenListV2Response, error) {
	result := &GetTaskChildrenListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetTaskChildrenListV2Uri(VERSION_V2)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskDetailV2
//
// PARAMS:
//   - request: the arguments to GetTaskDetailV2
//
// RETURNS:
//   - GetTaskDetailV2Response: The return type of the GetTaskDetailV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTaskDetailV2(request *GetTaskDetailV2Request) (*GetTaskDetailV2Response, error) {
	result := &GetTaskDetailV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTaskDetailV2Uri(VERSION_V2)).
		WithQueryParamFilter("namespace", util.StringValue(request.Namespace)).
		WithQueryParamFilter("dagId", util.StringValue(request.DagId)).
		WithQueryParamFilter("taskId", util.StringValue(request.TaskId)).
		WithQueryParamFilter("ignoreChildren", util.StringValue(request.IgnoreChildren)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTemplateDetailV2
//
// PARAMS:
//   - request: the arguments to GetTemplateDetailV2
//
// RETURNS:
//   - GetTemplateDetailV2Response: The return type of the GetTemplateDetailV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTemplateDetailV2(request *GetTemplateDetailV2Request) (*GetTemplateDetailV2Response, error) {
	result := &GetTemplateDetailV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTemplateDetailV2Uri(VERSION_V2)).
		WithQueryParamFilter("namespace", util.StringValue(request.Namespace)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTemplateListV2
//
// PARAMS:
//   - request: the arguments to GetTemplateListV2
//
// RETURNS:
//   - GetTemplateListV2Response: The return type of the GetTemplateListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTemplateListV2(request *GetTemplateListV2Request) (*GetTemplateListV2Response, error) {
	result := &GetTemplateListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetTemplateListV2Uri(VERSION_V2)).
		WithQueryParamFilter("locale", util.StringValue(request.Locale)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTemplateV2
//
// PARAMS:
//   - request: the arguments to UpdateTemplateV2
//
// RETURNS:
//   - UpdateTemplateV2Response: The return type of the UpdateTemplateV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateTemplateV2(request *UpdateTemplateV2Request) (*UpdateTemplateV2Response, error) {
	result := &UpdateTemplateV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateTemplateV2Uri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
