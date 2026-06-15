package bcm

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// CreateAlarmTemplate
//
// PARAMS:
//   - request: the arguments to CreateAlarmTemplate
//
// RETURNS:
//   - CreateAlarmTemplateResponse: The return type of the CreateAlarmTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAlarmTemplate(request *CreateAlarmTemplateRequest) (*CreateAlarmTemplateResponse, error) {
	result := &CreateAlarmTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAlarmTemplateUri()).
		WithQueryParamFilter("action", "CreateAlarmTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateInstanceGroup
//
// PARAMS:
//   - request: the arguments to CreateInstanceGroup
//
// RETURNS:
//   - CreateInstanceGroupResponse: The return type of the CreateInstanceGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceGroup(request *CreateInstanceGroupRequest) (*CreateInstanceGroupResponse, error) {
	result := &CreateInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceGroupUri()).
		WithQueryParamFilter("action", "CreateInstanceGroup").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateNotifyTemplate
//
// PARAMS:
//   - request: the arguments to CreateNotifyTemplate
//
// RETURNS:
//   - CreateNotifyTemplateResponse: The return type of the CreateNotifyTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateNotifyTemplate(request *CreateNotifyTemplateRequest) (*CreateNotifyTemplateResponse, error) {
	result := &CreateNotifyTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateNotifyTemplateUri()).
		WithQueryParamFilter("action", "CreateNotifyTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAlarmTemplates
//
// PARAMS:
//   - request: the arguments to DeleteAlarmTemplates
//
// RETURNS:
//   - DeleteAlarmTemplatesResponse: The return type of the DeleteAlarmTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAlarmTemplates(request *DeleteAlarmTemplatesRequest) (*DeleteAlarmTemplatesResponse, error) {
	result := &DeleteAlarmTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAlarmTemplatesUri()).
		WithQueryParamFilter("action", "DeleteAlarmTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteInstanceGroup
//
// PARAMS:
//   - request: the arguments to DeleteInstanceGroup
//
// RETURNS:
//   - DeleteInstanceGroupResponse: The return type of the DeleteInstanceGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceGroup(request *DeleteInstanceGroupRequest) (*DeleteInstanceGroupResponse, error) {
	result := &DeleteInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteInstanceGroupUri()).
		WithQueryParamFilter("action", "DeleteInstanceGroup").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteInstanceGroupInstances
//
// PARAMS:
//   - request: the arguments to DeleteInstanceGroupInstances
//
// RETURNS:
//   - DeleteInstanceGroupInstancesResponse: The return type of the DeleteInstanceGroupInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceGroupInstances(request *DeleteInstanceGroupInstancesRequest) (*DeleteInstanceGroupInstancesResponse, error) {
	result := &DeleteInstanceGroupInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteInstanceGroupInstancesUri()).
		WithQueryParamFilter("action", "DeleteInstanceGroupInstances").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteNotifyTemplate
//
// PARAMS:
//   - request: the arguments to DeleteNotifyTemplate
//
// RETURNS:
//   - DeleteNotifyTemplateResponse: The return type of the DeleteNotifyTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteNotifyTemplate(request *DeleteNotifyTemplateRequest) (*DeleteNotifyTemplateResponse, error) {
	result := &DeleteNotifyTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteNotifyTemplateUri()).
		WithQueryParamFilter("action", "DeleteNotifyTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmTemplate
//
// PARAMS:
//   - request: the arguments to DescribeAlarmTemplate
//
// RETURNS:
//   - DescribeAlarmTemplateResponse: The return type of the DescribeAlarmTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmTemplate(request *DescribeAlarmTemplateRequest) (*DescribeAlarmTemplateResponse, error) {
	result := &DescribeAlarmTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmTemplateUri()).
		WithQueryParamFilter("action", "DescribeAlarmTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmTemplates
//
// PARAMS:
//   - request: the arguments to DescribeAlarmTemplates
//
// RETURNS:
//   - DescribeAlarmTemplatesResponse: The return type of the DescribeAlarmTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmTemplates(request *DescribeAlarmTemplatesRequest) (*DescribeAlarmTemplatesResponse, error) {
	result := &DescribeAlarmTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmTemplatesUri()).
		WithQueryParamFilter("action", "DescribeAlarmTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeInstanceGroup
//
// PARAMS:
//   - request: the arguments to DescribeInstanceGroup
//
// RETURNS:
//   - DescribeInstanceGroupResponse: The return type of the DescribeInstanceGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeInstanceGroup(request *DescribeInstanceGroupRequest) (*DescribeInstanceGroupResponse, error) {
	result := &DescribeInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeInstanceGroupUri()).
		WithQueryParamFilter("action", "DescribeInstanceGroup").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeInstanceGroups
//
// PARAMS:
//   - request: the arguments to DescribeInstanceGroups
//
// RETURNS:
//   - DescribeInstanceGroupsResponse: The return type of the DescribeInstanceGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeInstanceGroups(request *DescribeInstanceGroupsRequest) (*DescribeInstanceGroupsResponse, error) {
	result := &DescribeInstanceGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeInstanceGroupsUri()).
		WithQueryParamFilter("action", "DescribeInstanceGroups").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeNotifyTemplate
//
// PARAMS:
//   - request: the arguments to DescribeNotifyTemplate
//
// RETURNS:
//   - DescribeNotifyTemplateResponse: The return type of the DescribeNotifyTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeNotifyTemplate(request *DescribeNotifyTemplateRequest) (*DescribeNotifyTemplateResponse, error) {
	result := &DescribeNotifyTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeNotifyTemplateUri()).
		WithQueryParamFilter("action", "DescribeNotifyTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeNotifyTemplates
//
// PARAMS:
//   - request: the arguments to DescribeNotifyTemplates
//
// RETURNS:
//   - DescribeNotifyTemplatesResponse: The return type of the DescribeNotifyTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeNotifyTemplates(request *DescribeNotifyTemplatesRequest) (*DescribeNotifyTemplatesResponse, error) {
	result := &DescribeNotifyTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeNotifyTemplatesUri()).
		WithQueryParamFilter("action", "DescribeNotifyTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeReceivers
//
// PARAMS:
//   - request: the arguments to DescribeReceivers
//
// RETURNS:
//   - DescribeReceiversResponse: The return type of the DescribeReceivers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeReceivers(request *DescribeReceiversRequest) (*DescribeReceiversResponse, error) {
	result := &DescribeReceiversResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeReceiversUri()).
		WithQueryParamFilter("action", "DescribeReceivers").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeSystemTemplateRules
//
// PARAMS:
//   - request: the arguments to DescribeSystemTemplateRules
//
// RETURNS:
//   - DescribeSystemTemplateRulesResponse: The return type of the DescribeSystemTemplateRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeSystemTemplateRules(request *DescribeSystemTemplateRulesRequest) (*DescribeSystemTemplateRulesResponse, error) {
	result := &DescribeSystemTemplateRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeSystemTemplateRulesUri()).
		WithQueryParamFilter("action", "DescribeSystemTemplateRules").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExportAlarmTemplates
//
// PARAMS:
//   - request: the arguments to ExportAlarmTemplates
//
// RETURNS:
//   - ExportAlarmTemplatesResponse: The return type of the ExportAlarmTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ExportAlarmTemplates(request *ExportAlarmTemplatesRequest) (*ExportAlarmTemplatesResponse, error) {
	result := &ExportAlarmTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getExportAlarmTemplatesUri()).
		WithQueryParamFilter("action", "ExportAlarmTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ImportAlarmTemplates
//
// PARAMS:
//   - request: the arguments to ImportAlarmTemplates
//
// RETURNS:
//   - ImportAlarmTemplatesResponse: The return type of the ImportAlarmTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImportAlarmTemplates(request *ImportAlarmTemplatesRequest) (*ImportAlarmTemplatesResponse, error) {
	result := &ImportAlarmTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImportAlarmTemplatesUri()).
		WithQueryParamFilter("action", "ImportAlarmTemplates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAlarmTemplate
//
// PARAMS:
//   - request: the arguments to UpdateAlarmTemplate
//
// RETURNS:
//   - UpdateAlarmTemplateResponse: The return type of the UpdateAlarmTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlarmTemplate(request *UpdateAlarmTemplateRequest) (*UpdateAlarmTemplateResponse, error) {
	result := &UpdateAlarmTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateAlarmTemplateUri()).
		WithQueryParamFilter("action", "UpdateAlarmTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateInstanceGroup
//
// PARAMS:
//   - request: the arguments to UpdateInstanceGroup
//
// RETURNS:
//   - UpdateInstanceGroupResponse: The return type of the UpdateInstanceGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceGroup(request *UpdateInstanceGroupRequest) (*UpdateInstanceGroupResponse, error) {
	result := &UpdateInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateInstanceGroupUri()).
		WithQueryParamFilter("action", "UpdateInstanceGroup").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateNotifyTemplate
//
// PARAMS:
//   - request: the arguments to UpdateNotifyTemplate
//
// RETURNS:
//   - UpdateNotifyTemplateResponse: The return type of the UpdateNotifyTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateNotifyTemplate(request *UpdateNotifyTemplateRequest) (*UpdateNotifyTemplateResponse, error) {
	result := &UpdateNotifyTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateNotifyTemplateUri()).
		WithQueryParamFilter("action", "UpdateNotifyTemplate").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
