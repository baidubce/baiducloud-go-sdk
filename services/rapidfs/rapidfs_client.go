package rapidfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V2 = "v2"
)

// AddCacheNodes
//
// PARAMS:
//   - request: the arguments to AddCacheNodes
//
// RETURNS:
//   - AddCacheNodesResponse: The return type of the AddCacheNodes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddCacheNodes(request *AddCacheNodesRequest) (*AddCacheNodesResponse, error) {
	result := &AddCacheNodesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "AddCacheNodes").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CancelCacheRuleJob
//
// PARAMS:
//   - request: the arguments to CancelCacheRuleJob
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CancelCacheRuleJob(request *CancelCacheRuleJobRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelCacheRuleJobUri(VERSION_V2)).
		WithQueryParamFilter("action", "CancelCacheRuleJob").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CancelMetaSyncJob
//
// PARAMS:
//   - request: the arguments to CancelMetaSyncJob
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CancelMetaSyncJob(request *CancelMetaSyncJobRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelMetaSyncJobUri(VERSION_V2)).
		WithQueryParamFilter("action", "CancelMetaSyncJob").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CheckBeforeAddCacheNodes
//
// PARAMS:
//   - request: the arguments to CheckBeforeAddCacheNodes
//
// RETURNS:
//   - CheckBeforeAddCacheNodesResponse: The return type of the CheckBeforeAddCacheNodes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CheckBeforeAddCacheNodes(request *CheckBeforeAddCacheNodesRequest) (*CheckBeforeAddCacheNodesResponse, error) {
	result := &CheckBeforeAddCacheNodesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCheckBeforeAddCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "CheckBeforeAddCacheNodes").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CheckBeforeCreateInstance
//
// PARAMS:
//   - request: the arguments to CheckBeforeCreateInstance
//
// RETURNS:
//   - CheckBeforeCreateInstanceResponse: The return type of the CheckBeforeCreateInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CheckBeforeCreateInstance(request *CheckBeforeCreateInstanceRequest) (*CheckBeforeCreateInstanceResponse, error) {
	result := &CheckBeforeCreateInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCheckBeforeCreateInstanceUri(VERSION_V2)).
		WithQueryParamFilter("action", "CheckBeforeCreateInstance").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAndAssignTag
//
// PARAMS:
//   - request: the arguments to CreateAndAssignTag
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAndAssignTag(request *CreateAndAssignTagRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAndAssignTagUri(VERSION_V2)).
		WithQueryParamFilter("action", "CreateAndAssignTag").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAuthGroup
//
// PARAMS:
//   - request: the arguments to CreateAuthGroup
//
// RETURNS:
//   - CreateAuthGroupResponse: The return type of the CreateAuthGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAuthGroup(request *CreateAuthGroupRequest) (*CreateAuthGroupResponse, error) {
	result := &CreateAuthGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAuthGroupUri(VERSION_V2)).
		WithQueryParamFilter("action", "CreateAuthGroup").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCacheRule
//
// PARAMS:
//   - request: the arguments to CreateCacheRule
//
// RETURNS:
//   - CreateCacheRuleResponse: The return type of the CreateCacheRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCacheRule(request *CreateCacheRuleRequest) (*CreateCacheRuleResponse, error) {
	result := &CreateCacheRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCacheRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "CreateCacheRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
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
//   - CreateInstanceResponse: The return type of the CreateInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(request *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	result := &CreateInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceUri(VERSION_V2)).
		WithQueryParamFilter("action", "CreateInstance").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateMetaSyncRule
//
// PARAMS:
//   - request: the arguments to CreateMetaSyncRule
//
// RETURNS:
//   - CreateMetaSyncRuleResponse: The return type of the CreateMetaSyncRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateMetaSyncRule(request *CreateMetaSyncRuleRequest) (*CreateMetaSyncRuleResponse, error) {
	result := &CreateMetaSyncRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateMetaSyncRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "CreateMetaSyncRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAuthGroup
//
// PARAMS:
//   - request: the arguments to DeleteAuthGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAuthGroup(request *DeleteAuthGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAuthGroupUri(VERSION_V2)).
		WithQueryParamFilter("action", "DeleteAuthGroup").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteCacheRule
//
// PARAMS:
//   - request: the arguments to DeleteCacheRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCacheRule(request *DeleteCacheRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteCacheRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "DeleteCacheRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteInstance
//
// PARAMS:
//   - request: the arguments to DeleteInstance
//
// RETURNS:
//   - DeleteInstanceResponse: The return type of the DeleteInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	result := &DeleteInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteInstanceUri(VERSION_V2)).
		WithQueryParamFilter("action", "DeleteInstance").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteMetaSyncRule
//
// PARAMS:
//   - request: the arguments to DeleteMetaSyncRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteMetaSyncRule(request *DeleteMetaSyncRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteMetaSyncRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "DeleteMetaSyncRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DescribeAihcResourcePools
//
// PARAMS:
//   - request: the arguments to DescribeAihcResourcePools
//
// RETURNS:
//   - DescribeAihcResourcePoolsResponse: The return type of the DescribeAihcResourcePools interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAihcResourcePools(request *DescribeAihcResourcePoolsRequest) (*DescribeAihcResourcePoolsResponse, error) {
	result := &DescribeAihcResourcePoolsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAihcResourcePoolsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeAihcResourcePools").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAllocatableDataSrcQuota
//
// PARAMS:
//   - request: the arguments to DescribeAllocatableDataSrcQuota
//
// RETURNS:
//   - DescribeAllocatableDataSrcQuotaResponse: The return type of the DescribeAllocatableDataSrcQuota interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAllocatableDataSrcQuota(request *DescribeAllocatableDataSrcQuotaRequest) (*DescribeAllocatableDataSrcQuotaResponse, error) {
	result := &DescribeAllocatableDataSrcQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAllocatableDataSrcQuotaUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeAllocatableDataSrcQuota").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAuthGroup
//
// PARAMS:
//   - request: the arguments to DescribeAuthGroup
//
// RETURNS:
//   - DescribeAuthGroupResponse: The return type of the DescribeAuthGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAuthGroup(request *DescribeAuthGroupRequest) (*DescribeAuthGroupResponse, error) {
	result := &DescribeAuthGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAuthGroupUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeAuthGroup").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAuthGroups
//
// PARAMS:
//   - request: the arguments to DescribeAuthGroups
//
// RETURNS:
//   - DescribeAuthGroupsResponse: The return type of the DescribeAuthGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAuthGroups(request *DescribeAuthGroupsRequest) (*DescribeAuthGroupsResponse, error) {
	result := &DescribeAuthGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAuthGroupsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeAuthGroups").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheDeployGroup
//
// PARAMS:
//   - request: the arguments to DescribeCacheDeployGroup
//
// RETURNS:
//   - DescribeCacheDeployGroupResponse: The return type of the DescribeCacheDeployGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheDeployGroup(request *DescribeCacheDeployGroupRequest) (*DescribeCacheDeployGroupResponse, error) {
	result := &DescribeCacheDeployGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheDeployGroupUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheDeployGroup").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheDeployGroups
//
// PARAMS:
//   - request: the arguments to DescribeCacheDeployGroups
//
// RETURNS:
//   - DescribeCacheDeployGroupsResponse: The return type of the DescribeCacheDeployGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheDeployGroups(request *DescribeCacheDeployGroupsRequest) (*DescribeCacheDeployGroupsResponse, error) {
	result := &DescribeCacheDeployGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheDeployGroupsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheDeployGroups").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheNode
//
// PARAMS:
//   - request: the arguments to DescribeCacheNode
//
// RETURNS:
//   - DescribeCacheNodeResponse: The return type of the DescribeCacheNode interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheNode(request *DescribeCacheNodeRequest) (*DescribeCacheNodeResponse, error) {
	result := &DescribeCacheNodeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheNodeUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheNode").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheNodeBccCandidates
//
// PARAMS:
//   - request: the arguments to DescribeCacheNodeBccCandidates
//
// RETURNS:
//   - DescribeCacheNodeBccCandidatesResponse: The return type of the DescribeCacheNodeBccCandidates interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheNodeBccCandidates(request *DescribeCacheNodeBccCandidatesRequest) (*DescribeCacheNodeBccCandidatesResponse, error) {
	result := &DescribeCacheNodeBccCandidatesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheNodeBccCandidatesUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheNodeBccCandidates").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheNodeQuota
//
// PARAMS:
//   - request: the arguments to DescribeCacheNodeQuota
//
// RETURNS:
//   - DescribeCacheNodeQuotaResponse: The return type of the DescribeCacheNodeQuota interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheNodeQuota(request *DescribeCacheNodeQuotaRequest) (*DescribeCacheNodeQuotaResponse, error) {
	result := &DescribeCacheNodeQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheNodeQuotaUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheNodeQuota").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheNodes
//
// PARAMS:
//   - request: the arguments to DescribeCacheNodes
//
// RETURNS:
//   - DescribeCacheNodesResponse: The return type of the DescribeCacheNodes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheNodes(request *DescribeCacheNodesRequest) (*DescribeCacheNodesResponse, error) {
	result := &DescribeCacheNodesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheNodes").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheRule
//
// PARAMS:
//   - request: the arguments to DescribeCacheRule
//
// RETURNS:
//   - DescribeCacheRuleResponse: The return type of the DescribeCacheRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheRule(request *DescribeCacheRuleRequest) (*DescribeCacheRuleResponse, error) {
	result := &DescribeCacheRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheRule").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheRuleJobs
//
// PARAMS:
//   - request: the arguments to DescribeCacheRuleJobs
//
// RETURNS:
//   - DescribeCacheRuleJobsResponse: The return type of the DescribeCacheRuleJobs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheRuleJobs(request *DescribeCacheRuleJobsRequest) (*DescribeCacheRuleJobsResponse, error) {
	result := &DescribeCacheRuleJobsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheRuleJobsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheRuleJobs").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCacheRules
//
// PARAMS:
//   - request: the arguments to DescribeCacheRules
//
// RETURNS:
//   - DescribeCacheRulesResponse: The return type of the DescribeCacheRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCacheRules(request *DescribeCacheRulesRequest) (*DescribeCacheRulesResponse, error) {
	result := &DescribeCacheRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCacheRulesUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCacheRules").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCceClusters
//
// PARAMS:
//   - request: the arguments to DescribeCceClusters
//
// RETURNS:
//   - DescribeCceClustersResponse: The return type of the DescribeCceClusters interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeCceClusters(request *DescribeCceClustersRequest) (*DescribeCceClustersResponse, error) {
	result := &DescribeCceClustersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeCceClustersUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeCceClusters").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDataSrc
//
// PARAMS:
//   - request: the arguments to DescribeDataSrc
//
// RETURNS:
//   - DescribeDataSrcResponse: The return type of the DescribeDataSrc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDataSrc(request *DescribeDataSrcRequest) (*DescribeDataSrcResponse, error) {
	result := &DescribeDataSrcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeDataSrcUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeDataSrc").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDataSrcs
//
// PARAMS:
//   - request: the arguments to DescribeDataSrcs
//
// RETURNS:
//   - DescribeDataSrcsResponse: The return type of the DescribeDataSrcs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDataSrcs(request *DescribeDataSrcsRequest) (*DescribeDataSrcsResponse, error) {
	result := &DescribeDataSrcsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeDataSrcsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeDataSrcs").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeInstance
//
// PARAMS:
//   - request: the arguments to DescribeInstance
//
// RETURNS:
//   - DescribeInstanceResponse: The return type of the DescribeInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (*DescribeInstanceResponse, error) {
	result := &DescribeInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeInstanceUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeInstance").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeInstances
//
// PARAMS:
//   - request: the arguments to DescribeInstances
//
// RETURNS:
//   - DescribeInstancesResponse: The return type of the DescribeInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (*DescribeInstancesResponse, error) {
	result := &DescribeInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeInstancesUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeInstances").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetaSyncJobs
//
// PARAMS:
//   - request: the arguments to DescribeMetaSyncJobs
//
// RETURNS:
//   - DescribeMetaSyncJobsResponse: The return type of the DescribeMetaSyncJobs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetaSyncJobs(request *DescribeMetaSyncJobsRequest) (*DescribeMetaSyncJobsResponse, error) {
	result := &DescribeMetaSyncJobsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetaSyncJobsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeMetaSyncJobs").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetaSyncRule
//
// PARAMS:
//   - request: the arguments to DescribeMetaSyncRule
//
// RETURNS:
//   - DescribeMetaSyncRuleResponse: The return type of the DescribeMetaSyncRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetaSyncRule(request *DescribeMetaSyncRuleRequest) (*DescribeMetaSyncRuleResponse, error) {
	result := &DescribeMetaSyncRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetaSyncRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeMetaSyncRule").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetaSyncRules
//
// PARAMS:
//   - request: the arguments to DescribeMetaSyncRules
//
// RETURNS:
//   - DescribeMetaSyncRulesResponse: The return type of the DescribeMetaSyncRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetaSyncRules(request *DescribeMetaSyncRulesRequest) (*DescribeMetaSyncRulesResponse, error) {
	result := &DescribeMetaSyncRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetaSyncRulesUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeMetaSyncRules").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeOrder
//
// PARAMS:
//   - request: the arguments to DescribeOrder
//
// RETURNS:
//   - DescribeOrderResponse: The return type of the DescribeOrder interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeOrder(request *DescribeOrderRequest) (*DescribeOrderResponse, error) {
	result := &DescribeOrderResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeOrderUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeOrder").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribePrice
//
// PARAMS:
//   - request: the arguments to DescribePrice
//
// RETURNS:
//   - DescribePriceResponse: The return type of the DescribePrice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribePrice(request *DescribePriceRequest) (*DescribePriceResponse, error) {
	result := &DescribePriceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribePriceUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribePrice").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeSpecs
//
// PARAMS:
//   - request: the arguments to DescribeSpecs
//
// RETURNS:
//   - DescribeSpecsResponse: The return type of the DescribeSpecs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeSpecs(request *DescribeSpecsRequest) (*DescribeSpecsResponse, error) {
	result := &DescribeSpecsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeSpecsUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeSpecs").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeToken
//
// PARAMS:
//   - request: the arguments to DescribeToken
//
// RETURNS:
//   - DescribeTokenResponse: The return type of the DescribeToken interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeToken(request *DescribeTokenRequest) (*DescribeTokenResponse, error) {
	result := &DescribeTokenResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeTokenUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeToken").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeZones
//
// PARAMS:
//   - request: the arguments to DescribeZones
//
// RETURNS:
//   - DescribeZonesResponse: The return type of the DescribeZones interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeZones() (*DescribeZonesResponse, error) {
	result := &DescribeZonesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeZonesUri(VERSION_V2)).
		WithQueryParamFilter("action", "DescribeZones").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DisableMetaSyncRule
//
// PARAMS:
//   - request: the arguments to DisableMetaSyncRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableMetaSyncRule(request *DisableMetaSyncRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDisableMetaSyncRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "DisableMetaSyncRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// EnableMetaSyncRule
//
// PARAMS:
//   - request: the arguments to EnableMetaSyncRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableMetaSyncRule(request *EnableMetaSyncRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEnableMetaSyncRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "EnableMetaSyncRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ExecuteCacheRuleJob
//
// PARAMS:
//   - request: the arguments to ExecuteCacheRuleJob
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ExecuteCacheRuleJob(request *ExecuteCacheRuleJobRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getExecuteCacheRuleJobUri(VERSION_V2)).
		WithQueryParamFilter("action", "ExecuteCacheRuleJob").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ExecuteMetaSyncJob
//
// PARAMS:
//   - request: the arguments to ExecuteMetaSyncJob
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ExecuteMetaSyncJob(request *ExecuteMetaSyncJobRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getExecuteMetaSyncJobUri(VERSION_V2)).
		WithQueryParamFilter("action", "ExecuteMetaSyncJob").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ImportDataSrc
//
// PARAMS:
//   - request: the arguments to ImportDataSrc
//
// RETURNS:
//   - ImportDataSrcResponse: The return type of the ImportDataSrc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImportDataSrc(request *ImportDataSrcRequest) (*ImportDataSrcResponse, error) {
	result := &ImportDataSrcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImportDataSrcUri(VERSION_V2)).
		WithQueryParamFilter("action", "ImportDataSrc").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyAuthGroup
//
// PARAMS:
//   - request: the arguments to ModifyAuthGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyAuthGroup(request *ModifyAuthGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyAuthGroupUri(VERSION_V2)).
		WithQueryParamFilter("action", "ModifyAuthGroup").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ModifyDataSrc
//
// PARAMS:
//   - request: the arguments to ModifyDataSrc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyDataSrc(request *ModifyDataSrcRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyDataSrcUri(VERSION_V2)).
		WithQueryParamFilter("action", "ModifyDataSrc").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ModifyMetaSyncRule
//
// PARAMS:
//   - request: the arguments to ModifyMetaSyncRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyMetaSyncRule(request *ModifyMetaSyncRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyMetaSyncRuleUri(VERSION_V2)).
		WithQueryParamFilter("action", "ModifyMetaSyncRule").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ModifyToken
//
// PARAMS:
//   - request: the arguments to ModifyToken
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyToken(request *ModifyTokenRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyTokenUri(VERSION_V2)).
		WithQueryParamFilter("action", "ModifyToken").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RemoveCacheNodes
//
// PARAMS:
//   - request: the arguments to RemoveCacheNodes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveCacheNodes(request *RemoveCacheNodesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "RemoveCacheNodes").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RemoveDataSrc
//
// PARAMS:
//   - request: the arguments to RemoveDataSrc
//
// RETURNS:
//   - RemoveDataSrcResponse: The return type of the RemoveDataSrc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RemoveDataSrc(request *RemoveDataSrcRequest) (*RemoveDataSrcResponse, error) {
	result := &RemoveDataSrcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveDataSrcUri(VERSION_V2)).
		WithQueryParamFilter("action", "RemoveDataSrc").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResizeInstance
//
// PARAMS:
//   - request: the arguments to ResizeInstance
//
// RETURNS:
//   - ResizeInstanceResponse: The return type of the ResizeInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeInstance(request *ResizeInstanceRequest) (*ResizeInstanceResponse, error) {
	result := &ResizeInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getResizeInstanceUri(VERSION_V2)).
		WithQueryParamFilter("action", "ResizeInstance").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RestartCacheNodes
//
// PARAMS:
//   - request: the arguments to RestartCacheNodes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestartCacheNodes(request *RestartCacheNodesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRestartCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "RestartCacheNodes").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// StartCacheNodes
//
// PARAMS:
//   - request: the arguments to StartCacheNodes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StartCacheNodes(request *StartCacheNodesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getStartCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "StartCacheNodes").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// StopCacheNodes
//
// PARAMS:
//   - request: the arguments to StopCacheNodes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StopCacheNodes(request *StopCacheNodesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getStopCacheNodesUri(VERSION_V2)).
		WithQueryParamFilter("action", "StopCacheNodes").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
