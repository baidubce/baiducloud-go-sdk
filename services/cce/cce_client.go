package cce

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// CreateAShrinkingNodeGroupTaskV2
//
// PARAMS:
//   - request: the arguments to CreateAShrinkingNodeGroupTaskV2
//
// RETURNS:
//   - CreateAShrinkingNodeGroupTaskV2Response: The return type of the CreateAShrinkingNodeGroupTaskV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAShrinkingNodeGroupTaskV2(request *CreateAShrinkingNodeGroupTaskV2Request) (*CreateAShrinkingNodeGroupTaskV2Response, error) {
	result := &CreateAShrinkingNodeGroupTaskV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCreateAShrinkingNodeGroupTaskV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAnAutoscalerV2
//
// PARAMS:
//   - request: the arguments to CreateAnAutoscalerV2
//
// RETURNS:
//   - CreateAnAutoscalerV2Response: The return type of the CreateAnAutoscalerV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAnAutoscalerV2(request *CreateAnAutoscalerV2Request) (*CreateAnAutoscalerV2Response, error) {
	result := &CreateAnAutoscalerV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAnAutoscalerV2Uri(util.StringValue(request.ClusterID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateExpansionNodeGroupTaskV2
//
// PARAMS:
//   - request: the arguments to CreateExpansionNodeGroupTaskV2
//
// RETURNS:
//   - CreateExpansionNodeGroupTaskV2Response: The return type of the CreateExpansionNodeGroupTaskV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateExpansionNodeGroupTaskV2(request *CreateExpansionNodeGroupTaskV2Request) (*CreateExpansionNodeGroupTaskV2Response, error) {
	result := &CreateExpansionNodeGroupTaskV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCreateExpansionNodeGroupTaskV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithQueryParamFilter("upToReplicas", util.Int32Value(request.UpToReplicas)).
		WithQueryParamFilter("upReplicas", util.Int32Value(request.UpReplicas)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateNodeGroupV2
//
// PARAMS:
//   - request: the arguments to CreateNodeGroupV2
//
// RETURNS:
//   - CreateNodeGroupV2Response: The return type of the CreateNodeGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateNodeGroupV2(request *CreateNodeGroupV2Request) (*CreateNodeGroupV2Response, error) {
	result := &CreateNodeGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateNodeGroupV2Uri(util.StringValue(request.ClusterID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateNodesClusterExpansionV2
//
// PARAMS:
//   - request: the arguments to CreateNodesClusterExpansionV2
//
// RETURNS:
//   - CreateNodesClusterExpansionV2Response: The return type of the CreateNodesClusterExpansionV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateNodesClusterExpansionV2(request *CreateNodesClusterExpansionV2Request) (*CreateNodesClusterExpansionV2Response, error) {
	result := &CreateNodesClusterExpansionV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateNodesClusterExpansionV2Uri(util.StringValue(request.ClusterID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteNodeGroupV2
//
// PARAMS:
//   - request: the arguments to DeleteNodeGroupV2
//
// RETURNS:
//   - DeleteNodeGroupV2Response: The return type of the DeleteNodeGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteNodeGroupV2(request *DeleteNodeGroupV2Request) (*DeleteNodeGroupV2Response, error) {
	result := &DeleteNodeGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteNodeGroupV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithQueryParamFilter("deleteInstances", util.BoolValue(request.DeleteInstances)).
		WithQueryParamFilter("releaseAllResource", util.BoolValue(request.ReleaseAllResource)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteNodesClusterScalingV2
//
// PARAMS:
//   - request: the arguments to DeleteNodesClusterScalingV2
//
// RETURNS:
//   - DeleteNodesClusterScalingV2Response: The return type of the DeleteNodesClusterScalingV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteNodesClusterScalingV2(request *DeleteNodesClusterScalingV2Request) (*DeleteNodesClusterScalingV2Response, error) {
	result := &DeleteNodesClusterScalingV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteNodesClusterScalingV2Uri(util.StringValue(request.ClusterID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodeDetailsV2
//
// PARAMS:
//   - request: the arguments to GetNodeDetailsV2
//
// RETURNS:
//   - GetNodeDetailsV2Response: The return type of the GetNodeDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetNodeDetailsV2(request *GetNodeDetailsV2Request) (*GetNodeDetailsV2Response, error) {
	result := &GetNodeDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetNodeDetailsV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodeGroupDetailsV2
//
// PARAMS:
//   - request: the arguments to GetNodeGroupDetailsV2
//
// RETURNS:
//   - GetNodeGroupDetailsV2Response: The return type of the GetNodeGroupDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetNodeGroupDetailsV2(request *GetNodeGroupDetailsV2Request) (*GetNodeGroupDetailsV2Response, error) {
	result := &GetNodeGroupDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetNodeGroupDetailsV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPackageListV2
//
// PARAMS:
//   - request: the arguments to GetPackageListV2
//
// RETURNS:
//   - GetPackageListV2Response: The return type of the GetPackageListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPackageListV2(request *GetPackageListV2Request) (*GetPackageListV2Response, error) {
	result := &GetPackageListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetPackageListV2Uri()).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskListV2
//
// PARAMS:
//   - request: the arguments to GetTaskListV2
//
// RETURNS:
//   - GetTaskListV2Response: The return type of the GetTaskListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTaskListV2(request *GetTaskListV2Request) (*GetTaskListV2Response, error) {
	result := &GetTaskListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTaskListV2Uri(util.StringValue(request.TaskType))).
		WithQueryParamFilter("targetID", util.StringValue(request.TargetID)).
		WithQueryParamFilter("operationType", util.StringValue(request.OperationType)).
		WithQueryParamFilter("phase", util.StringValue(request.Phase)).
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

// GetTheListOfClusterNodeGroupsV2
//
// PARAMS:
//   - request: the arguments to GetTheListOfClusterNodeGroupsV2
//
// RETURNS:
//   - GetTheListOfClusterNodeGroupsV2Response: The return type of the GetTheListOfClusterNodeGroupsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTheListOfClusterNodeGroupsV2(request *GetTheListOfClusterNodeGroupsV2Request) (*GetTheListOfClusterNodeGroupsV2Response, error) {
	result := &GetTheListOfClusterNodeGroupsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTheListOfClusterNodeGroupsV2Uri(util.StringValue(request.ClusterID))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("autoscalerEnabled", util.StringValue(request.AutoscalerEnabled)).
		WithQueryParamFilter("chargingType", util.StringValue(request.ChargingType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTheListOfClusterNodesV2
//
// PARAMS:
//   - request: the arguments to GetTheListOfClusterNodesV2
//
// RETURNS:
//   - GetTheListOfClusterNodesV2Response: The return type of the GetTheListOfClusterNodesV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTheListOfClusterNodesV2(request *GetTheListOfClusterNodesV2Request) (*GetTheListOfClusterNodesV2Response, error) {
	result := &GetTheListOfClusterNodesV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTheListOfClusterNodesV2Uri(util.StringValue(request.ClusterID))).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
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

// ModifyIGAutoScaler
//
// PARAMS:
//   - request: the arguments to ModifyIGAutoScaler
//
// RETURNS:
//   - ModifyIGAutoScalerResponse: The return type of the ModifyIGAutoScaler interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyIGAutoScaler(request *ModifyIGAutoScalerRequest) (*ModifyIGAutoScalerResponse, error) {
	result := &ModifyIGAutoScalerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyIGAutoScalerUri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyNodeGroupNodeShrinkProtectionStatusV2
//
// PARAMS:
//   - request: the arguments to ModifyNodeGroupNodeShrinkProtectionStatusV2
//
// RETURNS:
//   - ModifyNodeGroupNodeShrinkProtectionStatusV2Response: The return type of the ModifyNodeGroupNodeShrinkProtectionStatusV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyNodeGroupNodeShrinkProtectionStatusV2(request *ModifyNodeGroupNodeShrinkProtectionStatusV2Request) (*ModifyNodeGroupNodeShrinkProtectionStatusV2Response, error) {
	result := &ModifyNodeGroupNodeShrinkProtectionStatusV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyNodeGroupNodeShrinkProtectionStatusV2Uri(util.StringValue(request.ClusterID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyTheNumberOfNodeReplicasInANodeGroupV2
//
// PARAMS:
//   - request: the arguments to ModifyTheNumberOfNodeReplicasInANodeGroupV2
//
// RETURNS:
//   - ModifyTheNumberOfNodeReplicasInANodeGroupV2Response: The return type of the ModifyTheNumberOfNodeReplicasInANodeGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyTheNumberOfNodeReplicasInANodeGroupV2(request *ModifyTheNumberOfNodeReplicasInANodeGroupV2Request) (*ModifyTheNumberOfNodeReplicasInANodeGroupV2Response, error) {
	result := &ModifyTheNumberOfNodeReplicasInANodeGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTheNumberOfNodeReplicasInANodeGroupV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MoveIntoAnExistingNodeV2
//
// PARAMS:
//   - request: the arguments to MoveIntoAnExistingNodeV2
//
// RETURNS:
//   - MoveIntoAnExistingNodeV2Response: The return type of the MoveIntoAnExistingNodeV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MoveIntoAnExistingNodeV2(request *MoveIntoAnExistingNodeV2Request) (*MoveIntoAnExistingNodeV2Response, error) {
	result := &MoveIntoAnExistingNodeV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMoveIntoAnExistingNodeV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheConfigurationOfAutoscalerV2
//
// PARAMS:
//   - request: the arguments to QueryTheConfigurationOfAutoscalerV2
//
// RETURNS:
//   - QueryTheConfigurationOfAutoscalerV2Response: The return type of the QueryTheConfigurationOfAutoscalerV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheConfigurationOfAutoscalerV2(request *QueryTheConfigurationOfAutoscalerV2Request) (*QueryTheConfigurationOfAutoscalerV2Response, error) {
	result := &QueryTheConfigurationOfAutoscalerV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheConfigurationOfAutoscalerV2Uri(util.StringValue(request.ClusterID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RetrieveTheNodeGroupNodeListV2
//
// PARAMS:
//   - request: the arguments to RetrieveTheNodeGroupNodeListV2
//
// RETURNS:
//   - RetrieveTheNodeGroupNodeListV2Response: The return type of the RetrieveTheNodeGroupNodeListV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveTheNodeGroupNodeListV2(request *RetrieveTheNodeGroupNodeListV2Request) (*RetrieveTheNodeGroupNodeListV2Response, error) {
	result := &RetrieveTheNodeGroupNodeListV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveTheNodeGroupNodeListV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceGroupID))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StepsToObtainNodeEventsV2
//
// PARAMS:
//   - request: the arguments to StepsToObtainNodeEventsV2
//
// RETURNS:
//   - StepsToObtainNodeEventsV2Response: The return type of the StepsToObtainNodeEventsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) StepsToObtainNodeEventsV2(request *StepsToObtainNodeEventsV2Request) (*StepsToObtainNodeEventsV2Response, error) {
	result := &StepsToObtainNodeEventsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getStepsToObtainNodeEventsV2Uri(util.StringValue(request.InstanceID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SynchronizeNodeMetadataV2
//
// PARAMS:
//   - request: the arguments to SynchronizeNodeMetadataV2
//
// RETURNS:
//   - SynchronizeNodeMetadataV2Response: The return type of the SynchronizeNodeMetadataV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SynchronizeNodeMetadataV2(request *SynchronizeNodeMetadataV2Request) (*SynchronizeNodeMetadataV2Response, error) {
	result := &SynchronizeNodeMetadataV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSynchronizeNodeMetadataV2Uri(util.StringValue(request.ClusterID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAutoscalerConfigurationV2
//
// PARAMS:
//   - request: the arguments to UpdateAutoscalerConfigurationV2
//
// RETURNS:
//   - UpdateAutoscalerConfigurationV2Response: The return type of the UpdateAutoscalerConfigurationV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAutoscalerConfigurationV2(request *UpdateAutoscalerConfigurationV2Request) (*UpdateAutoscalerConfigurationV2Response, error) {
	result := &UpdateAutoscalerConfigurationV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAutoscalerConfigurationV2Uri(util.StringValue(request.ClusterID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateNodeAttributesV2
//
// PARAMS:
//   - request: the arguments to UpdateNodeAttributesV2
//
// RETURNS:
//   - UpdateNodeAttributesV2Response: The return type of the UpdateNodeAttributesV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateNodeAttributesV2(request *UpdateNodeAttributesV2Request) (*UpdateNodeAttributesV2Response, error) {
	result := &UpdateNodeAttributesV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateNodeAttributesV2Uri(util.StringValue(request.ClusterID), util.StringValue(request.InstanceID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ViewTaskDetailsV2
//
// PARAMS:
//   - request: the arguments to ViewTaskDetailsV2
//
// RETURNS:
//   - ViewTaskDetailsV2Response: The return type of the ViewTaskDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ViewTaskDetailsV2(request *ViewTaskDetailsV2Request) (*ViewTaskDetailsV2Response, error) {
	result := &ViewTaskDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getViewTaskDetailsV2Uri(util.StringValue(request.TaskType), util.StringValue(request.TaskID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
