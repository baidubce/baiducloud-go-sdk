package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// BatchConfigurationOfInstances
//
// PARAMS:
//   - request: the arguments to BatchConfigurationOfInstances
//
// RETURNS:
//   - BatchConfigurationOfInstancesResponse: The return type of the BatchConfigurationOfInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchConfigurationOfInstances(request *BatchConfigurationOfInstancesRequest) (*BatchConfigurationOfInstancesResponse, error) {
	result := &BatchConfigurationOfInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBatchConfigurationOfInstancesUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindingTags
//
// PARAMS:
//   - request: the arguments to BindingTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindingTags(request *BindingTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBindingTagsUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// ClusterStatusCheck
//
// PARAMS:
//   - request: the arguments to ClusterStatusCheck
//
// RETURNS:
//   - ClusterStatusCheckResponse: The return type of the ClusterStatusCheck interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ClusterStatusCheck(request *ClusterStatusCheckRequest) (*ClusterStatusCheckResponse, error) {
	result := &ClusterStatusCheckResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterStatusCheckUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAPrimaryInstance
//
// PARAMS:
//   - request: the arguments to CreateAPrimaryInstance
//
// RETURNS:
//   - CreateAPrimaryInstanceResponse: The return type of the CreateAPrimaryInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAPrimaryInstance(request *CreateAPrimaryInstanceRequest) (*CreateAPrimaryInstanceResponse, error) {
	result := &CreateAPrimaryInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAPrimaryInstanceUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAProxyInstance
//
// PARAMS:
//   - request: the arguments to CreateAProxyInstance
//
// RETURNS:
//   - CreateAProxyInstanceResponse: The return type of the CreateAProxyInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAProxyInstance(request *CreateAProxyInstanceRequest) (*CreateAProxyInstanceResponse, error) {
	result := &CreateAProxyInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAProxyInstanceUri(VERSION_V1)).
		WithQueryParam("rdsproxy", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateReadOnlyInstance
//
// PARAMS:
//   - request: the arguments to CreateReadOnlyInstance
//
// RETURNS:
//   - CreateReadOnlyInstanceResponse: The return type of the CreateReadOnlyInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateReadOnlyInstance(request *CreateReadOnlyInstanceRequest) (*CreateReadOnlyInstanceResponse, error) {
	result := &CreateReadOnlyInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateReadOnlyInstanceUri(VERSION_V1)).
		WithQueryParam("readReplica", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DialTestInterface
//
// PARAMS:
//   - request: the arguments to DialTestInterface
//
// RETURNS:
//   - DialTestInterfaceResponse: The return type of the DialTestInterface interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DialTestInterface(request *DialTestInterfaceRequest) (*DialTestInterfaceResponse, error) {
	result := &DialTestInterfaceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDialTestInterfaceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DoesTheInstanceSupportEnablingAutomaticExpansion
//
// PARAMS:
//   - request: the arguments to DoesTheInstanceSupportEnablingAutomaticExpansion
//
// RETURNS:
//   - DoesTheInstanceSupportEnablingAutomaticExpansionResponse: The return type of the DoesTheInstanceSupportEnablingAutomaticExpansion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DoesTheInstanceSupportEnablingAutomaticExpansion(request *DoesTheInstanceSupportEnablingAutomaticExpansionRequest) (*DoesTheInstanceSupportEnablingAutomaticExpansionResponse, error) {
	result := &DoesTheInstanceSupportEnablingAutomaticExpansionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDoesTheInstanceSupportEnablingAutomaticExpansionUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InstanceEnablesAutomaticRenewal
//
// PARAMS:
//   - request: the arguments to InstanceEnablesAutomaticRenewal
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) InstanceEnablesAutomaticRenewal(request *InstanceEnablesAutomaticRenewalRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceEnablesAutomaticRenewalUri(VERSION_V1)).
		WithQueryParam("autoRenew", "").
		WithBody(request).
		Do()
}

// QueryInstanceList
//
// PARAMS:
//   - request: the arguments to QueryInstanceList
//
// RETURNS:
//   - QueryInstanceListResponse: The return type of the QueryInstanceList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryInstanceList(request *QueryInstanceListRequest) (*QueryInstanceListResponse, error) {
	result := &QueryInstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryInstanceListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("vnetIp", util.StringValue(request.VnetIp)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySubnetList
//
// PARAMS:
//   - request: the arguments to QuerySubnetList
//
// RETURNS:
//   - QuerySubnetListResponse: The return type of the QuerySubnetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySubnetList(request *QuerySubnetListRequest) (*QuerySubnetListResponse, error) {
	result := &QuerySubnetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySubnetListUri(VERSION_V1)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTaskList
//
// PARAMS:
//   - request: the arguments to QueryTaskList
//
// RETURNS:
//   - QueryTaskListResponse: The return type of the QueryTaskList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTaskList(request *QueryTaskListRequest) (*QueryTaskListResponse, error) {
	result := &QueryTaskListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryTaskListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheDetailsOfTheSpecifiedInstance
//
// PARAMS:
//   - request: the arguments to QueryTheDetailsOfTheSpecifiedInstance
//
// RETURNS:
//   - QueryTheDetailsOfTheSpecifiedInstanceResponse: The return type of the QueryTheDetailsOfTheSpecifiedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheDetailsOfTheSpecifiedInstance(request *QueryTheDetailsOfTheSpecifiedInstanceRequest) (*QueryTheDetailsOfTheSpecifiedInstanceResponse, error) {
	result := &QueryTheDetailsOfTheSpecifiedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheDetailsOfTheSpecifiedInstanceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfAvailableZones
//
// PARAMS:
//   - request: the arguments to QueryTheListOfAvailableZones
//
// RETURNS:
//   - QueryTheListOfAvailableZonesResponse: The return type of the QueryTheListOfAvailableZones interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfAvailableZones() (*QueryTheListOfAvailableZonesResponse, error) {
	result := &QueryTheListOfAvailableZonesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfAvailableZonesUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseInstance
//
// PARAMS:
//   - request: the arguments to ReleaseInstance
//
// RETURNS:
//   - ReleaseInstanceResponse: The return type of the ReleaseInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ReleaseInstance(request *ReleaseInstanceRequest) (*ReleaseInstanceResponse, error) {
	result := &ReleaseInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseInstanceUri(VERSION_V1)).
		WithQueryParamFilter("instanceIds", util.StringValue(request.InstanceIds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RenewalExample
//
// PARAMS:
//   - request: the arguments to RenewalExample
//
// RETURNS:
//   - RenewalExampleResponse: The return type of the RenewalExample interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RenewalExample(request *RenewalExampleRequest) (*RenewalExampleResponse, error) {
	result := &RenewalExampleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRenewalExampleUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RestartTheInstance
//
// PARAMS:
//   - request: the arguments to RestartTheInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestartTheInstance(request *RestartTheInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRestartTheInstanceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("reboot", "").
		WithBody(request).
		Do()
}

// RetrieveInstanceAutoScalingConfig
//
// PARAMS:
//   - request: the arguments to RetrieveInstanceAutoScalingConfig
//
// RETURNS:
//   - RetrieveInstanceAutoScalingConfigResponse: The return type of the RetrieveInstanceAutoScalingConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveInstanceAutoScalingConfig(request *RetrieveInstanceAutoScalingConfigRequest) (*RetrieveInstanceAutoScalingConfigResponse, error) {
	result := &RetrieveInstanceAutoScalingConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveInstanceAutoScalingConfigUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TransformationExample
//
// PARAMS:
//   - request: the arguments to TransformationExample
//
// RETURNS:
//   - TransformationExampleResponse: The return type of the TransformationExample interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TransformationExample(request *TransformationExampleRequest) (*TransformationExampleResponse, error) {
	result := &TransformationExampleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getTransformationExampleUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAutomaticStorageExpansionConfiguration
//
// PARAMS:
//   - request: the arguments to UpdateAutomaticStorageExpansionConfiguration
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAutomaticStorageExpansionConfiguration(request *UpdateAutomaticStorageExpansionConfigurationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAutomaticStorageExpansionConfigurationUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.Action))).
		WithBody(request).
		Do()
}

// UpdateInstanceConnectionInformation
//
// PARAMS:
//   - request: the arguments to UpdateInstanceConnectionInformation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceConnectionInformation(request *UpdateInstanceConnectionInformationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceConnectionInformationUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyEndpoint", "").
		WithBody(request).
		Do()
}

// UpdateInstanceName
//
// PARAMS:
//   - request: the arguments to UpdateInstanceName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceNameUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("rename", "").
		WithBody(request).
		Do()
}

// UpdateInstancePublicNetworkAccessStatus
//
// PARAMS:
//   - request: the arguments to UpdateInstancePublicNetworkAccessStatus
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstancePublicNetworkAccessStatus(request *UpdateInstancePublicNetworkAccessStatusRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstancePublicNetworkAccessStatusUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyPublicAccess", "").
		WithBody(request).
		Do()
}

// UpdateInstanceSynchronizationMode
//
// PARAMS:
//   - request: the arguments to UpdateInstanceSynchronizationMode
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceSynchronizationMode(request *UpdateInstanceSynchronizationModeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceSynchronizationModeUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifySyncMode", "").
		WithBody(request).
		Do()
}

// UpdateTimeWindow
//
// PARAMS:
//   - request: the arguments to UpdateTimeWindow
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTimeWindow(request *UpdateTimeWindowRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateTimeWindowUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}
