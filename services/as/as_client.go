package as

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// AdjustNumV2
//
// PARAMS:
//   - request: the arguments to AdjustNumV2
//
// RETURNS:
//   - AdjustNumV2Response: The return type of the AdjustNumV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AdjustNumV2(request *AdjustNumV2Request) (*AdjustNumV2Response, error) {
	result := &AdjustNumV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAdjustNumV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("adjustNode", util.StringValue(request.AdjustNode)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AttachNodeV2
//
// PARAMS:
//   - request: the arguments to AttachNodeV2
//
// RETURNS:
//   - AttachNodeV2Response: The return type of the AttachNodeV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AttachNodeV2(request *AttachNodeV2Request) (*AttachNodeV2Response, error) {
	result := &AttachNodeV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAttachNodeV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("attachNode", util.StringValue(request.AttachNode)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAsGroupV2
//
// PARAMS:
//   - request: the arguments to CreateAsGroupV2
//
// RETURNS:
//   - CreateAsGroupV2Response: The return type of the CreateAsGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAsGroupV2(request *CreateAsGroupV2Request) (*CreateAsGroupV2Response, error) {
	result := &CreateAsGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAsGroupV2Uri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRuleV2
//
// PARAMS:
//   - request: the arguments to CreateRuleV2
//
// RETURNS:
//   - CreateRuleV2Response: The return type of the CreateRuleV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateRuleV2(request *CreateRuleV2Request) (*CreateRuleV2Response, error) {
	result := &CreateRuleV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRuleV2Uri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAsGroupV2
//
// PARAMS:
//   - request: the arguments to DeleteAsGroupV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAsGroupV2(request *DeleteAsGroupV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAsGroupV2Uri()).
		WithBody(request).
		Do()
}

// DeleteRuleV2
//
// PARAMS:
//   - request: the arguments to DeleteRuleV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRuleV2(request *DeleteRuleV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRuleV2Uri(util.StringValue(request.RuleId))).
		Do()
}

// DetachNodeV2
//
// PARAMS:
//   - request: the arguments to DetachNodeV2
//
// RETURNS:
//   - DetachNodeV2Response: The return type of the DetachNodeV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DetachNodeV2(request *DetachNodeV2Request) (*DetachNodeV2Response, error) {
	result := &DetachNodeV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDetachNodeV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("detachNode", util.StringValue(request.DetachNode)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecRuleV2
//
// PARAMS:
//   - request: the arguments to ExecRuleV2
//
// RETURNS:
//   - ExecRuleV2Response: The return type of the ExecRuleV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ExecRuleV2(request *ExecRuleV2Request) (*ExecRuleV2Response, error) {
	result := &ExecRuleV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getExecRuleV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("execRule", util.StringValue(request.ExecRule)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAsGroupV2
//
// PARAMS:
//   - request: the arguments to GetAsGroupV2
//
// RETURNS:
//   - GetAsGroupV2Response: The return type of the GetAsGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAsGroupV2(request *GetAsGroupV2Request) (*GetAsGroupV2Response, error) {
	result := &GetAsGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAsGroupV2Uri(util.StringValue(request.GroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRuleV2
//
// PARAMS:
//   - request: the arguments to GetRuleV2
//
// RETURNS:
//   - GetRuleV2Response: The return type of the GetRuleV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRuleV2(request *GetRuleV2Request) (*GetRuleV2Response, error) {
	result := &GetRuleV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRuleV2Uri(util.StringValue(request.RuleId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAsGroupV2
//
// PARAMS:
//   - request: the arguments to ListAsGroupV2
//
// RETURNS:
//   - ListAsGroupV2Response: The return type of the ListAsGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAsGroupV2(request *ListAsGroupV2Request) (*ListAsGroupV2Response, error) {
	result := &ListAsGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAsGroupV2Uri()).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("subKeywordType", util.StringValue(request.SubKeywordType)).
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

// ListAsNodeV2
//
// PARAMS:
//   - request: the arguments to ListAsNodeV2
//
// RETURNS:
//   - ListAsNodeV2Response: The return type of the ListAsNodeV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAsNodeV2(request *ListAsNodeV2Request) (*ListAsNodeV2Response, error) {
	result := &ListAsNodeV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAsNodeV2Uri()).
		WithQueryParamFilter("groupid", util.StringValue(request.Groupid)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
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

// ListRuleV2
//
// PARAMS:
//   - request: the arguments to ListRuleV2
//
// RETURNS:
//   - ListRuleV2Response: The return type of the ListRuleV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRuleV2(request *ListRuleV2Request) (*ListRuleV2Response, error) {
	result := &ListRuleV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRuleV2Uri()).
		WithQueryParamFilter("groupid", util.StringValue(request.Groupid)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
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

// ListTaskV2
//
// PARAMS:
//   - request: the arguments to ListTaskV2
//
// RETURNS:
//   - ListTaskV2Response: The return type of the ListTaskV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTaskV2(request *ListTaskV2Request) (*ListTaskV2Response, error) {
	result := &ListTaskV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTaskV2Uri()).
		WithQueryParamFilter("groupid", util.StringValue(request.Groupid)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("startTime", util.StringValue(request.StartTime)).
		WithQueryParamFilter("endTime", util.StringValue(request.EndTime)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ScalingDownV2
//
// PARAMS:
//   - request: the arguments to ScalingDownV2
//
// RETURNS:
//   - ScalingDownV2Response: The return type of the ScalingDownV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ScalingDownV2(request *ScalingDownV2Request) (*ScalingDownV2Response, error) {
	result := &ScalingDownV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getScalingDownV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("scalingDown", util.StringValue(request.ScalingDown)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ScalingUpV2
//
// PARAMS:
//   - request: the arguments to ScalingUpV2
//
// RETURNS:
//   - ScalingUpV2Response: The return type of the ScalingUpV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ScalingUpV2(request *ScalingUpV2Request) (*ScalingUpV2Response, error) {
	result := &ScalingUpV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getScalingUpV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("scalingUp", util.StringValue(request.ScalingUp)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIsManagedV2
//
// PARAMS:
//   - request: the arguments to UpdateIsManagedV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIsManagedV2(request *UpdateIsManagedV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIsManagedV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("updateIsManaged", util.StringValue(request.UpdateIsManaged)).
		WithBody(request).
		Do()
}

// UpdateProtectV2
//
// PARAMS:
//   - request: the arguments to UpdateProtectV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateProtectV2(request *UpdateProtectV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateProtectV2Uri(util.StringValue(request.GroupId))).
		WithQueryParamFilter("updateProtect", util.StringValue(request.UpdateProtect)).
		WithBody(request).
		Do()
}
