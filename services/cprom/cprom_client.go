package cprom

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// BindCluster
//
// PARAMS:
//   - request: the arguments to BindCluster
//
// RETURNS:
//   - BindClusterResponse: The return type of the BindCluster interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BindCluster(request *BindClusterRequest) (*BindClusterResponse, error) {
	result := &BindClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindClusterUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ClaimAlertEvent
//
// PARAMS:
//   - request: the arguments to ClaimAlertEvent
//
// RETURNS:
//   - ClaimAlertEventResponse: The return type of the ClaimAlertEvent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ClaimAlertEvent(request *ClaimAlertEventRequest) (*ClaimAlertEventResponse, error) {
	result := &ClaimAlertEventResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClaimAlertEventUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAlert
//
// PARAMS:
//   - request: the arguments to CreateAlert
//
// RETURNS:
//   - CreateAlertResponse: The return type of the CreateAlert interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAlert(request *CreateAlertRequest) (*CreateAlertResponse, error) {
	result := &CreateAlertResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAlertUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCustomScrapeTask
//
// PARAMS:
//   - request: the arguments to CreateCustomScrapeTask
//
// RETURNS:
//   - CreateCustomScrapeTaskResponse: The return type of the CreateCustomScrapeTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCustomScrapeTask(request *CreateCustomScrapeTaskRequest) (*CreateCustomScrapeTaskResponse, error) {
	result := &CreateCustomScrapeTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCustomScrapeTaskUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateInstance
//
// PARAMS:
//   - request: the arguments to CreateInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(request *CreateInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceUri()).
		WithBody(request).
		Do()
}

// CreateNotificationPolicy
//
// PARAMS:
//   - request: the arguments to CreateNotificationPolicy
//
// RETURNS:
//   - CreateNotificationPolicyResponse: The return type of the CreateNotificationPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateNotificationPolicy(request *CreateNotificationPolicyRequest) (*CreateNotificationPolicyResponse, error) {
	result := &CreateNotificationPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateNotificationPolicyUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePodmonitor
//
// PARAMS:
//   - request: the arguments to CreatePodmonitor
//
// RETURNS:
//   - CreatePodmonitorResponse: The return type of the CreatePodmonitor interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePodmonitor(request *CreatePodmonitorRequest) (*CreatePodmonitorResponse, error) {
	result := &CreatePodmonitorResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePodmonitorUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateServiceMonitor
//
// PARAMS:
//   - request: the arguments to CreateServiceMonitor
//
// RETURNS:
//   - CreateServiceMonitorResponse: The return type of the CreateServiceMonitor interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateServiceMonitor(request *CreateServiceMonitorRequest) (*CreateServiceMonitorResponse, error) {
	result := &CreateServiceMonitorResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateServiceMonitorUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAlert
//
// PARAMS:
//   - request: the arguments to DeleteAlert
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAlert(request *DeleteAlertRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAlertUri(util.StringValue(request.AlertingRuleId))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		Do()
}

// DeleteCustomScrapeTask
//
// PARAMS:
//   - request: the arguments to DeleteCustomScrapeTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCustomScrapeTask(request *DeleteCustomScrapeTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteCustomScrapeTaskUri(util.StringValue(request.ScrapeJobId))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		Do()
}

// DeleteInstance
//
// PARAMS:
//   - request: the arguments to DeleteInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteInstanceUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("deleteGrafana", util.BoolValue(request.DeleteGrafana)).
		Do()
}

// DeleteNotificationPolicy
//
// PARAMS:
//   - request: the arguments to DeleteNotificationPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteNotificationPolicy(request *DeleteNotificationPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteNotificationPolicyUri(util.StringValue(request.NotifyRuleId))).
		Do()
}

// DeletePodmonitor
//
// PARAMS:
//   - request: the arguments to DeletePodmonitor
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePodmonitor(request *DeletePodmonitorRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePodmonitorUri(util.StringValue(request.PodMonitorName))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		Do()
}

// DeleteServiceMonitor
//
// PARAMS:
//   - request: the arguments to DeleteServiceMonitor
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteServiceMonitor(request *DeleteServiceMonitorRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteServiceMonitorUri(util.StringValue(request.ServiceMonitorName))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		Do()
}

// GenerateInstanceToken
//
// PARAMS:
//   - request: the arguments to GenerateInstanceToken
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GenerateInstanceToken(request *GenerateInstanceTokenRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGenerateInstanceTokenUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("Action", util.StringValue(request.Action)).
		WithBody(request).
		Do()
}

// GetAlertDetail
//
// PARAMS:
//   - request: the arguments to GetAlertDetail
//
// RETURNS:
//   - GetAlertDetailResponse: The return type of the GetAlertDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAlertDetail(request *GetAlertDetailRequest) (*GetAlertDetailResponse, error) {
	result := &GetAlertDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAlertDetailUri(util.StringValue(request.AlertingRuleId))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAlertEventDetail
//
// PARAMS:
//   - request: the arguments to GetAlertEventDetail
//
// RETURNS:
//   - GetAlertEventDetailResponse: The return type of the GetAlertEventDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAlertEventDetail(request *GetAlertEventDetailRequest) (*GetAlertEventDetailResponse, error) {
	result := &GetAlertEventDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAlertEventDetailUri(util.StringValue(request.EventId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetClusterBindStatus
//
// PARAMS:
//   - request: the arguments to GetClusterBindStatus
//
// RETURNS:
//   - GetClusterBindStatusResponse: The return type of the GetClusterBindStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetClusterBindStatus(request *GetClusterBindStatusRequest) (*GetClusterBindStatusResponse, error) {
	result := &GetClusterBindStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetClusterBindStatusUri()).
		WithQueryParamFilter("clusterId", util.StringValue(request.ClusterId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNotificationPolicy
//
// PARAMS:
//   - request: the arguments to GetNotificationPolicy
//
// RETURNS:
//   - GetNotificationPolicyResponse: The return type of the GetNotificationPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetNotificationPolicy(request *GetNotificationPolicyRequest) (*GetNotificationPolicyResponse, error) {
	result := &GetNotificationPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetNotificationPolicyUri(util.StringValue(request.NotifyRuleId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPodMonitorDetail
//
// PARAMS:
//   - request: the arguments to GetPodMonitorDetail
//
// RETURNS:
//   - GetPodMonitorDetailResponse: The return type of the GetPodMonitorDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPodMonitorDetail(request *GetPodMonitorDetailRequest) (*GetPodMonitorDetailResponse, error) {
	result := &GetPodMonitorDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPodMonitorDetailUri(util.StringValue(request.PodMonitorName))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetServiceMonitorDetail
//
// PARAMS:
//   - request: the arguments to GetServiceMonitorDetail
//
// RETURNS:
//   - GetServiceMonitorDetailResponse: The return type of the GetServiceMonitorDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetServiceMonitorDetail(request *GetServiceMonitorDetailRequest) (*GetServiceMonitorDetailResponse, error) {
	result := &GetServiceMonitorDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetServiceMonitorDetailUri(util.StringValue(request.ServiceMonitorName))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlertEvents
//
// PARAMS:
//   - request: the arguments to ListAlertEvents
//
// RETURNS:
//   - ListAlertEventsResponse: The return type of the ListAlertEvents interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlertEvents(request *ListAlertEventsRequest) (*ListAlertEventsResponse, error) {
	result := &ListAlertEventsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAlertEventsUri()).
		WithQueryParamFilter("startTime", util.Int32Value(request.StartTime)).
		WithQueryParamFilter("endTime", util.Int32Value(request.EndTime)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("monitorInstanceId", util.StringValue(request.MonitorInstanceId)).
		WithQueryParamFilter("alertingRuleId", util.StringValue(request.AlertingRuleId)).
		WithQueryParamFilter("alertingRuleName", util.StringValue(request.AlertingRuleName)).
		WithQueryParamFilter("notifyRuleId", util.StringValue(request.NotifyRuleId)).
		WithQueryParamFilter("notifyRuleName", util.StringValue(request.NotifyRuleName)).
		WithQueryParamFilter("severity", util.StringValue(request.Severity)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("expr", util.StringValue(request.Expr)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("alarmTags", util.StringValue(request.AlarmTags)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlertTemplates
//
// PARAMS:
//   - request: the arguments to ListAlertTemplates
//
// RETURNS:
//   - ListAlertTemplatesResponse: The return type of the ListAlertTemplates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlertTemplates() (*ListAlertTemplatesResponse, error) {
	result := &ListAlertTemplatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAlertTemplatesUri()).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAlerts
//
// PARAMS:
//   - request: the arguments to ListAlerts
//
// RETURNS:
//   - ListAlertsResponse: The return type of the ListAlerts interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAlerts(request *ListAlertsRequest) (*ListAlertsResponse, error) {
	result := &ListAlertsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAlertsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListBindableCloudProducts
//
// PARAMS:
//   - request: the arguments to ListBindableCloudProducts
//
// RETURNS:
//   - ListBindableCloudProductsResponse: The return type of the ListBindableCloudProducts interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListBindableCloudProducts() (*ListBindableCloudProductsResponse, error) {
	result := &ListBindableCloudProductsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListBindableCloudProductsUri()).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstances
//
// PARAMS:
//   - request: the arguments to ListInstances
//
// RETURNS:
//   - ListInstancesResponse: The return type of the ListInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(request *ListInstancesRequest) (*ListInstancesResponse, error) {
	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstancesUri()).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("phase", util.StringValue(request.Phase)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListNotificationPolicies
//
// PARAMS:
//   - request: the arguments to ListNotificationPolicies
//
// RETURNS:
//   - ListNotificationPoliciesResponse: The return type of the ListNotificationPolicies interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListNotificationPolicies(request *ListNotificationPoliciesRequest) (*ListNotificationPoliciesResponse, error) {
	result := &ListNotificationPoliciesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListNotificationPoliciesUri()).
		WithQueryParamFilter("pageNo", "1").
		WithQueryParamFilter("pageSize", "10").
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPodMonitors
//
// PARAMS:
//   - request: the arguments to ListPodMonitors
//
// RETURNS:
//   - ListPodMonitorsResponse: The return type of the ListPodMonitors interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPodMonitors(request *ListPodMonitorsRequest) (*ListPodMonitorsResponse, error) {
	result := &ListPodMonitorsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPodMonitorsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRelatedCloudProducts
//
// PARAMS:
//   - request: the arguments to ListRelatedCloudProducts
//
// RETURNS:
//   - ListRelatedCloudProductsResponse: The return type of the ListRelatedCloudProducts interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRelatedCloudProducts(request *ListRelatedCloudProductsRequest) (*ListRelatedCloudProductsResponse, error) {
	result := &ListRelatedCloudProductsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRelatedCloudProductsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListServiceMonitors
//
// PARAMS:
//   - request: the arguments to ListServiceMonitors
//
// RETURNS:
//   - ListServiceMonitorsResponse: The return type of the ListServiceMonitors interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListServiceMonitors(request *ListServiceMonitorsRequest) (*ListServiceMonitorsResponse, error) {
	result := &ListServiceMonitorsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListServiceMonitorsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoteRead
//
// PARAMS:
//   - request: the arguments to RemoteRead
//
// RETURNS:
//   - RemoteReadResponse: The return type of the RemoteRead interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RemoteRead(request *RemoteReadRequest) (*RemoteReadResponse, error) {
	result := &RemoteReadResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoteReadUri(util.StringValue(request.RemoteReadUrl))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoteWrite
//
// PARAMS:
//   - request: the arguments to RemoteWrite
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoteWrite(request *RemoteWriteRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoteWriteUri(util.StringValue(request.RemoteWriteUrl))).
		WithHeaderFilter("Content-Type", util.StringValue(request.ContentType)).
		WithHeaderFilter("Content-Encoding", util.StringValue(request.ContentEncoding)).
		WithHeaderFilter("InstanceId", util.StringValue(request.InstanceId)).
		WithHeaderFilter("Authorization", util.StringValue(request.Authorization)).
		WithStreamField(request.Body).
		Do()
}

// TogglePodMonitorService
//
// PARAMS:
//   - request: the arguments to TogglePodMonitorService
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) TogglePodMonitorService(request *TogglePodMonitorServiceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getTogglePodMonitorServiceUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		Do()
}

// ToggleServiceMonitorService
//
// PARAMS:
//   - request: the arguments to ToggleServiceMonitorService
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ToggleServiceMonitorService(request *ToggleServiceMonitorServiceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getToggleServiceMonitorServiceUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		Do()
}

// UpdateAlert
//
// PARAMS:
//   - request: the arguments to UpdateAlert
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAlert(request *UpdateAlertRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAlertUri(util.StringValue(request.AlertingRuleId))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithBody(request).
		Do()
}

// UpdateNotificationPolicy
//
// PARAMS:
//   - request: the arguments to UpdateNotificationPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateNotificationPolicy(request *UpdateNotificationPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateNotificationPolicyUri(util.StringValue(request.NotifyRuleId))).
		WithBody(request).
		Do()
}

// UpdatePodMonitor
//
// PARAMS:
//   - request: the arguments to UpdatePodMonitor
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePodMonitor(request *UpdatePodMonitorRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePodMonitorUri(util.StringValue(request.PodMonitorName))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithBody(request).
		Do()
}

// UpdateRelatedCloudProducts
//
// PARAMS:
//   - request: the arguments to UpdateRelatedCloudProducts
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRelatedCloudProducts(request *UpdateRelatedCloudProductsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRelatedCloudProductsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithBody(request).
		Do()
}

// UpdateServiceMonitor
//
// PARAMS:
//   - request: the arguments to UpdateServiceMonitor
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateServiceMonitor(request *UpdateServiceMonitorRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateServiceMonitorUri(util.StringValue(request.ServiceMonitorName))).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("agentId", util.StringValue(request.AgentId)).
		WithBody(request).
		Do()
}
