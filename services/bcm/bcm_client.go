package bcm

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// AddAlarmPolicyActions
//
// PARAMS:
//   - request: the arguments to AddAlarmPolicyActions
//
// RETURNS:
//   - AddAlarmPolicyActionsResponse: The return type of the AddAlarmPolicyActions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddAlarmPolicyActions(request *AddAlarmPolicyActionsRequest) (*AddAlarmPolicyActionsResponse, error) {
	result := &AddAlarmPolicyActionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddAlarmPolicyActionsUri()).
		WithQueryParamFilter("action", "AddAlarmPolicyActions").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAlarmMasking
//
// PARAMS:
//   - request: the arguments to CreateAlarmMasking
//
// RETURNS:
//   - CreateAlarmMaskingResponse: The return type of the CreateAlarmMasking interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAlarmMasking(request *CreateAlarmMaskingRequest) (*CreateAlarmMaskingResponse, error) {
	result := &CreateAlarmMaskingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAlarmMaskingUri()).
		WithQueryParamFilter("action", "CreateAlarmMasking").
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
		WithURL(getCreateAlarmPolicyUri()).
		WithQueryParamFilter("action", "CreateAlarmPolicy").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

// DeleteAlarmMaskings
//
// PARAMS:
//   - request: the arguments to DeleteAlarmMaskings
//
// RETURNS:
//   - DeleteAlarmMaskingsResponse: The return type of the DeleteAlarmMaskings interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAlarmMaskings(request *DeleteAlarmMaskingsRequest) (*DeleteAlarmMaskingsResponse, error) {
	result := &DeleteAlarmMaskingsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAlarmMaskingsUri()).
		WithQueryParamFilter("action", "DeleteAlarmMaskings").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAlarmPolicies
//
// PARAMS:
//   - request: the arguments to DeleteAlarmPolicies
//
// RETURNS:
//   - DeleteAlarmPoliciesResponse: The return type of the DeleteAlarmPolicies interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAlarmPolicies(request *DeleteAlarmPoliciesRequest) (*DeleteAlarmPoliciesResponse, error) {
	result := &DeleteAlarmPoliciesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAlarmPoliciesUri()).
		WithQueryParamFilter("action", "DeleteAlarmPolicies").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAlarmPolicyActions
//
// PARAMS:
//   - request: the arguments to DeleteAlarmPolicyActions
//
// RETURNS:
//   - DeleteAlarmPolicyActionsResponse: The return type of the DeleteAlarmPolicyActions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAlarmPolicyActions(request *DeleteAlarmPolicyActionsRequest) (*DeleteAlarmPolicyActionsResponse, error) {
	result := &DeleteAlarmPolicyActionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAlarmPolicyActionsUri()).
		WithQueryParamFilter("action", "DeleteAlarmPolicyActions").
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

// DescribeAlarm
//
// PARAMS:
//   - request: the arguments to DescribeAlarm
//
// RETURNS:
//   - DescribeAlarmResponse: The return type of the DescribeAlarm interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarm(request *DescribeAlarmRequest) (*DescribeAlarmResponse, error) {
	result := &DescribeAlarmResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmUri()).
		WithQueryParamFilter("action", "DescribeAlarm").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmMasking
//
// PARAMS:
//   - request: the arguments to DescribeAlarmMasking
//
// RETURNS:
//   - DescribeAlarmMaskingResponse: The return type of the DescribeAlarmMasking interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmMasking(request *DescribeAlarmMaskingRequest) (*DescribeAlarmMaskingResponse, error) {
	result := &DescribeAlarmMaskingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmMaskingUri()).
		WithQueryParamFilter("action", "DescribeAlarmMasking").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmMaskings
//
// PARAMS:
//   - request: the arguments to DescribeAlarmMaskings
//
// RETURNS:
//   - DescribeAlarmMaskingsResponse: The return type of the DescribeAlarmMaskings interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmMaskings(request *DescribeAlarmMaskingsRequest) (*DescribeAlarmMaskingsResponse, error) {
	result := &DescribeAlarmMaskingsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmMaskingsUri()).
		WithQueryParamFilter("action", "DescribeAlarmMaskings").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAlarmPolicies
//
// PARAMS:
//   - request: the arguments to DescribeAlarmPolicies
//
// RETURNS:
//   - DescribeAlarmPoliciesResponse: The return type of the DescribeAlarmPolicies interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarmPolicies(request *DescribeAlarmPoliciesRequest) (*DescribeAlarmPoliciesResponse, error) {
	result := &DescribeAlarmPoliciesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmPoliciesUri()).
		WithQueryParamFilter("action", "DescribeAlarmPolicies").
		WithBody(request).
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
		WithMethod(http.POST).
		WithURL(getDescribeAlarmPolicyUri()).
		WithQueryParamFilter("action", "DescribeAlarmPolicy").
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

// DescribeAlarms
//
// PARAMS:
//   - request: the arguments to DescribeAlarms
//
// RETURNS:
//   - DescribeAlarmsResponse: The return type of the DescribeAlarms interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (*DescribeAlarmsResponse, error) {
	result := &DescribeAlarmsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAlarmsUri()).
		WithQueryParamFilter("action", "DescribeAlarms").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDimensionValues
//
// PARAMS:
//   - request: the arguments to DescribeDimensionValues
//
// RETURNS:
//   - DescribeDimensionValuesResponse: The return type of the DescribeDimensionValues interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDimensionValues(request *DescribeDimensionValuesRequest) (*DescribeDimensionValuesResponse, error) {
	result := &DescribeDimensionValuesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeDimensionValuesUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeDimensionValues").
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

// DescribeMetricData
//
// PARAMS:
//   - request: the arguments to DescribeMetricData
//
// RETURNS:
//   - DescribeMetricDataResponse: The return type of the DescribeMetricData interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetricData(request *DescribeMetricDataRequest) (*DescribeMetricDataResponse, error) {
	result := &DescribeMetricDataResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetricDataUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeMetricData").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetricDataLatest
//
// PARAMS:
//   - request: the arguments to DescribeMetricDataLatest
//
// RETURNS:
//   - DescribeMetricDataLatestResponse: The return type of the DescribeMetricDataLatest interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetricDataLatest(request *DescribeMetricDataLatestRequest) (*DescribeMetricDataLatestResponse, error) {
	result := &DescribeMetricDataLatestResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetricDataLatestUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeMetricDataLatest").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetricDataLatestTop
//
// PARAMS:
//   - request: the arguments to DescribeMetricDataLatestTop
//
// RETURNS:
//   - DescribeMetricDataLatestTopResponse: The return type of the DescribeMetricDataLatestTop interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetricDataLatestTop(request *DescribeMetricDataLatestTopRequest) (*DescribeMetricDataLatestTopResponse, error) {
	result := &DescribeMetricDataLatestTopResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetricDataLatestTopUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeMetricDataLatestTop").
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

// UpdateAlarmMasking
//
// PARAMS:
//   - request: the arguments to UpdateAlarmMasking
//
// RETURNS:
//   - UpdateAlarmMaskingResponse: The return type of the UpdateAlarmMasking interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlarmMasking(request *UpdateAlarmMaskingRequest) (*UpdateAlarmMaskingResponse, error) {
	result := &UpdateAlarmMaskingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateAlarmMaskingUri()).
		WithQueryParamFilter("action", "UpdateAlarmMasking").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAlarmMaskingStates
//
// PARAMS:
//   - request: the arguments to UpdateAlarmMaskingStates
//
// RETURNS:
//   - UpdateAlarmMaskingStatesResponse: The return type of the UpdateAlarmMaskingStates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlarmMaskingStates(request *UpdateAlarmMaskingStatesRequest) (*UpdateAlarmMaskingStatesResponse, error) {
	result := &UpdateAlarmMaskingStatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateAlarmMaskingStatesUri()).
		WithQueryParamFilter("action", "UpdateAlarmMaskingStates").
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
		WithMethod(http.POST).
		WithURL(getUpdateAlarmPolicyUri()).
		WithQueryParamFilter("action", "UpdateAlarmPolicy").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAlarmPolicyNotifyEnabled
//
// PARAMS:
//   - request: the arguments to UpdateAlarmPolicyNotifyEnabled
//
// RETURNS:
//   - UpdateAlarmPolicyNotifyEnabledResponse: The return type of the UpdateAlarmPolicyNotifyEnabled interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlarmPolicyNotifyEnabled(request *UpdateAlarmPolicyNotifyEnabledRequest) (*UpdateAlarmPolicyNotifyEnabledResponse, error) {
	result := &UpdateAlarmPolicyNotifyEnabledResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateAlarmPolicyNotifyEnabledUri()).
		WithQueryParamFilter("action", "UpdateAlarmPolicyNotifyEnabled").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAlarmPolicyState
//
// PARAMS:
//   - request: the arguments to UpdateAlarmPolicyState
//
// RETURNS:
//   - UpdateAlarmPolicyStateResponse: The return type of the UpdateAlarmPolicyState interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlarmPolicyState(request *UpdateAlarmPolicyStateRequest) (*UpdateAlarmPolicyStateResponse, error) {
	result := &UpdateAlarmPolicyStateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateAlarmPolicyStateUri()).
		WithQueryParamFilter("action", "UpdateAlarmPolicyState").
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
