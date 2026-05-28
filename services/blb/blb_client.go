package blb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddAppBlbServerGroupRs
//
// PARAMS:
//   - request: the arguments to AddAppBlbServerGroupRs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddAppBlbServerGroupRs(request *AddAppBlbServerGroupRsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddAppBlbServerGroupRsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddBlbServer
//
// PARAMS:
//   - request: the arguments to AddBlbServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddBlbServer(request *AddBlbServerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddBlbServerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BillingChangeCancelToPostBlb
//
// PARAMS:
//   - request: the arguments to BillingChangeCancelToPostBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BillingChangeCancelToPostBlb(request *BillingChangeCancelToPostBlbRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBillingChangeCancelToPostBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("action", "CANCEL_TO_POSTPAY").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// BillingChangePostToPreBlb
//
// PARAMS:
//   - request: the arguments to BillingChangePostToPreBlb
//
// RETURNS:
//   - BillingChangePostToPreBlbResponse: The return type of the BillingChangePostToPreBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BillingChangePostToPreBlb(request *BillingChangePostToPreBlbRequest) (*BillingChangePostToPreBlbResponse, error) {
	result := &BillingChangePostToPreBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBillingChangePostToPreBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("action", "TO_PREPAY").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BillingChangePreToPostBlb
//
// PARAMS:
//   - request: the arguments to BillingChangePreToPostBlb
//
// RETURNS:
//   - BillingChangePreToPostBlbResponse: The return type of the BillingChangePreToPostBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BillingChangePreToPostBlb(request *BillingChangePreToPostBlbRequest) (*BillingChangePreToPostBlbResponse, error) {
	result := &BillingChangePreToPostBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBillingChangePreToPostBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("action", "TO_POSTPAY").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindBlbEnterpriseSecurityGroup
//
// PARAMS:
//   - request: the arguments to BindBlbEnterpriseSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindBlbEnterpriseSecurityGroup(request *BindBlbEnterpriseSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindBlbEnterpriseSecurityGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindBlbSecurityGroup
//
// PARAMS:
//   - request: the arguments to BindBlbSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindBlbSecurityGroup(request *BindBlbSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindBlbSecurityGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BlbInquiry
//
// PARAMS:
//   - request: the arguments to BlbInquiry
//
// RETURNS:
//   - BlbInquiryResponse: The return type of the BlbInquiry interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BlbInquiry(request *BlbInquiryRequest) (*BlbInquiryResponse, error) {
	result := &BlbInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBlbInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAppBlb
//
// PARAMS:
//   - request: the arguments to CreateAppBlb
//
// RETURNS:
//   - CreateAppBlbResponse: The return type of the CreateAppBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlb(request *CreateAppBlbRequest) (*CreateAppBlbResponse, error) {
	result := &CreateAppBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAppBlbHttpListener
//
// PARAMS:
//   - request: the arguments to CreateAppBlbHttpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbHttpListener(request *CreateAppBlbHttpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbHttpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAppBlbHttpsListener
//
// PARAMS:
//   - request: the arguments to CreateAppBlbHttpsListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbHttpsListener(request *CreateAppBlbHttpsListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbHttpsListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAppBlbPolicy
//
// PARAMS:
//   - request: the arguments to CreateAppBlbPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbPolicy(request *CreateAppBlbPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbPolicyUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAppBlbServerGroup
//
// PARAMS:
//   - request: the arguments to CreateAppBlbServerGroup
//
// RETURNS:
//   - CreateAppBlbServerGroupResponse: The return type of the CreateAppBlbServerGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbServerGroup(request *CreateAppBlbServerGroupRequest) (*CreateAppBlbServerGroupResponse, error) {
	result := &CreateAppBlbServerGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbServerGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAppBlbServerGroupPort
//
// PARAMS:
//   - request: the arguments to CreateAppBlbServerGroupPort
//
// RETURNS:
//   - CreateAppBlbServerGroupPortResponse: The return type of the CreateAppBlbServerGroupPort interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbServerGroupPort(request *CreateAppBlbServerGroupPortRequest) (*CreateAppBlbServerGroupPortResponse, error) {
	result := &CreateAppBlbServerGroupPortResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbServerGroupPortUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAppBlbSslListener
//
// PARAMS:
//   - request: the arguments to CreateAppBlbSslListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbSslListener(request *CreateAppBlbSslListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbSslListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAppBlbTcpListener
//
// PARAMS:
//   - request: the arguments to CreateAppBlbTcpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbTcpListener(request *CreateAppBlbTcpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbTcpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAppBlbUdpListener
//
// PARAMS:
//   - request: the arguments to CreateAppBlbUdpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAppBlbUdpListener(request *CreateAppBlbUdpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAppBlbUdpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateBlb
//
// PARAMS:
//   - request: the arguments to CreateBlb
//
// RETURNS:
//   - CreateBlbResponse: The return type of the CreateBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateBlb(request *CreateBlbRequest) (*CreateBlbResponse, error) {
	result := &CreateBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBlbUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateBlbHttpListener
//
// PARAMS:
//   - request: the arguments to CreateBlbHttpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateBlbHttpListener(request *CreateBlbHttpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBlbHttpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithBody(request).
		Do()
}

// CreateBlbHttpsListener
//
// PARAMS:
//   - request: the arguments to CreateBlbHttpsListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateBlbHttpsListener(request *CreateBlbHttpsListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBlbHttpsListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateBlbSslListener
//
// PARAMS:
//   - request: the arguments to CreateBlbSslListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateBlbSslListener(request *CreateBlbSslListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBlbSslListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateBlbTcpListener
//
// PARAMS:
//   - request: the arguments to CreateBlbTcpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateBlbTcpListener(request *CreateBlbTcpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBlbTcpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithBody(request).
		Do()
}

// CreateBlbUdpListener
//
// PARAMS:
//   - request: the arguments to CreateBlbUdpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateBlbUdpListener(request *CreateBlbUdpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBlbUdpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteAppBlbListener
//
// PARAMS:
//   - request: the arguments to DeleteAppBlbListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAppBlbListener(request *DeleteAppBlbListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteAppBlbListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("batchdelete", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteAppBlbPolicy
//
// PARAMS:
//   - request: the arguments to DeleteAppBlbPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAppBlbPolicy(request *DeleteAppBlbPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteAppBlbPolicyUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("batchdelete", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteAppBlbServerGroup
//
// PARAMS:
//   - request: the arguments to DeleteAppBlbServerGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAppBlbServerGroup(request *DeleteAppBlbServerGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteAppBlbServerGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("delete", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteAppBlbServerGroupPort
//
// PARAMS:
//   - request: the arguments to DeleteAppBlbServerGroupPort
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAppBlbServerGroupPort(request *DeleteAppBlbServerGroupPortRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteAppBlbServerGroupPortUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("batchdelete", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteAppBlbServerGroupRs
//
// PARAMS:
//   - request: the arguments to DeleteAppBlbServerGroupRs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAppBlbServerGroupRs(request *DeleteAppBlbServerGroupRsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteAppBlbServerGroupRsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("batchdelete", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteBlbListener
//
// PARAMS:
//   - request: the arguments to DeleteBlbListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteBlbListener(request *DeleteBlbListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteBlbListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("batchdelete", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteBlbServer
//
// PARAMS:
//   - request: the arguments to DeleteBlbServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteBlbServer(request *DeleteBlbServerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteBlbServerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DescribeAppBlb
//
// PARAMS:
//   - request: the arguments to DescribeAppBlb
//
// RETURNS:
//   - DescribeAppBlbResponse: The return type of the DescribeAppBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlb(request *DescribeAppBlbRequest) (*DescribeAppBlbResponse, error) {
	result := &DescribeAppBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbHttpListener
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbHttpListener
//
// RETURNS:
//   - DescribeAppBlbHttpListenerResponse: The return type of the DescribeAppBlbHttpListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbHttpListener(request *DescribeAppBlbHttpListenerRequest) (*DescribeAppBlbHttpListenerResponse, error) {
	result := &DescribeAppBlbHttpListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbHttpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbHttpsListener
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbHttpsListener
//
// RETURNS:
//   - DescribeAppBlbHttpsListenerResponse: The return type of the DescribeAppBlbHttpsListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbHttpsListener(request *DescribeAppBlbHttpsListenerRequest) (*DescribeAppBlbHttpsListenerResponse, error) {
	result := &DescribeAppBlbHttpsListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbHttpsListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbListener
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbListener
//
// RETURNS:
//   - DescribeAppBlbListenerResponse: The return type of the DescribeAppBlbListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbListener(request *DescribeAppBlbListenerRequest) (*DescribeAppBlbListenerResponse, error) {
	result := &DescribeAppBlbListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbPolicy
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbPolicy
//
// RETURNS:
//   - DescribeAppBlbPolicyResponse: The return type of the DescribeAppBlbPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbPolicy(request *DescribeAppBlbPolicyRequest) (*DescribeAppBlbPolicyResponse, error) {
	result := &DescribeAppBlbPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbPolicyUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("port", util.Int32Value(request.Port)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbServerGroup
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbServerGroup
//
// RETURNS:
//   - DescribeAppBlbServerGroupResponse: The return type of the DescribeAppBlbServerGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbServerGroup(request *DescribeAppBlbServerGroupRequest) (*DescribeAppBlbServerGroupResponse, error) {
	result := &DescribeAppBlbServerGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbServerGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("exactlyMatch", util.BoolValue(request.ExactlyMatch)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbServerGroupMountRs
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbServerGroupMountRs
//
// RETURNS:
//   - DescribeAppBlbServerGroupMountRsResponse: The return type of the DescribeAppBlbServerGroupMountRs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbServerGroupMountRs(request *DescribeAppBlbServerGroupMountRsRequest) (*DescribeAppBlbServerGroupMountRsResponse, error) {
	result := &DescribeAppBlbServerGroupMountRsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbServerGroupMountRsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("sgId", util.StringValue(request.SgId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbServerGroupRs
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbServerGroupRs
//
// RETURNS:
//   - DescribeAppBlbServerGroupRsResponse: The return type of the DescribeAppBlbServerGroupRs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbServerGroupRs(request *DescribeAppBlbServerGroupRsRequest) (*DescribeAppBlbServerGroupRsResponse, error) {
	result := &DescribeAppBlbServerGroupRsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbServerGroupRsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("sgId", util.StringValue(request.SgId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbServerGroupUnmountRs
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbServerGroupUnmountRs
//
// RETURNS:
//   - DescribeAppBlbServerGroupUnmountRsResponse: The return type of the DescribeAppBlbServerGroupUnmountRs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbServerGroupUnmountRs(request *DescribeAppBlbServerGroupUnmountRsRequest) (*DescribeAppBlbServerGroupUnmountRsResponse, error) {
	result := &DescribeAppBlbServerGroupUnmountRsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbServerGroupUnmountRsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("sgId", util.StringValue(request.SgId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbSslListener
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbSslListener
//
// RETURNS:
//   - DescribeAppBlbSslListenerResponse: The return type of the DescribeAppBlbSslListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbSslListener(request *DescribeAppBlbSslListenerRequest) (*DescribeAppBlbSslListenerResponse, error) {
	result := &DescribeAppBlbSslListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbSslListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbTcpListener
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbTcpListener
//
// RETURNS:
//   - DescribeAppBlbTcpListenerResponse: The return type of the DescribeAppBlbTcpListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbTcpListener(request *DescribeAppBlbTcpListenerRequest) (*DescribeAppBlbTcpListenerResponse, error) {
	result := &DescribeAppBlbTcpListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbTcpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbUdpListener
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbUdpListener
//
// RETURNS:
//   - DescribeAppBlbUdpListenerResponse: The return type of the DescribeAppBlbUdpListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbUdpListener(request *DescribeAppBlbUdpListenerRequest) (*DescribeAppBlbUdpListenerResponse, error) {
	result := &DescribeAppBlbUdpListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbUdpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeAppBlbs
//
// PARAMS:
//   - request: the arguments to DescribeAppBlbs
//
// RETURNS:
//   - DescribeAppBlbsResponse: The return type of the DescribeAppBlbs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAppBlbs(request *DescribeAppBlbsRequest) (*DescribeAppBlbsResponse, error) {
	result := &DescribeAppBlbsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeAppBlbsUri(VERSION_V1)).
		WithQueryParamFilter("address", util.StringValue(request.Address)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("blbId", util.StringValue(request.BlbId)).
		WithQueryParamFilter("bccId", util.StringValue(request.BccId)).
		WithQueryParamFilter("exactlyMatch", util.BoolValue(request.ExactlyMatch)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlb
//
// PARAMS:
//   - request: the arguments to DescribeBlb
//
// RETURNS:
//   - DescribeBlbResponse: The return type of the DescribeBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlb(request *DescribeBlbRequest) (*DescribeBlbResponse, error) {
	result := &DescribeBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbEnterpriseSecurityGroups
//
// PARAMS:
//   - request: the arguments to DescribeBlbEnterpriseSecurityGroups
//
// RETURNS:
//   - DescribeBlbEnterpriseSecurityGroupsResponse: The return type of the DescribeBlbEnterpriseSecurityGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbEnterpriseSecurityGroups(request *DescribeBlbEnterpriseSecurityGroupsRequest) (*DescribeBlbEnterpriseSecurityGroupsResponse, error) {
	result := &DescribeBlbEnterpriseSecurityGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbEnterpriseSecurityGroupsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbHttpListener
//
// PARAMS:
//   - request: the arguments to DescribeBlbHttpListener
//
// RETURNS:
//   - DescribeBlbHttpListenerResponse: The return type of the DescribeBlbHttpListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbHttpListener(request *DescribeBlbHttpListenerRequest) (*DescribeBlbHttpListenerResponse, error) {
	result := &DescribeBlbHttpListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbHttpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbHttpsListener
//
// PARAMS:
//   - request: the arguments to DescribeBlbHttpsListener
//
// RETURNS:
//   - DescribeBlbHttpsListenerResponse: The return type of the DescribeBlbHttpsListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbHttpsListener(request *DescribeBlbHttpsListenerRequest) (*DescribeBlbHttpsListenerResponse, error) {
	result := &DescribeBlbHttpsListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbHttpsListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbListener
//
// PARAMS:
//   - request: the arguments to DescribeBlbListener
//
// RETURNS:
//   - DescribeBlbListenerResponse: The return type of the DescribeBlbListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbListener(request *DescribeBlbListenerRequest) (*DescribeBlbListenerResponse, error) {
	result := &DescribeBlbListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbSecurityGroups
//
// PARAMS:
//   - request: the arguments to DescribeBlbSecurityGroups
//
// RETURNS:
//   - DescribeBlbSecurityGroupsResponse: The return type of the DescribeBlbSecurityGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbSecurityGroups(request *DescribeBlbSecurityGroupsRequest) (*DescribeBlbSecurityGroupsResponse, error) {
	result := &DescribeBlbSecurityGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbSecurityGroupsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbServerHealth
//
// PARAMS:
//   - request: the arguments to DescribeBlbServerHealth
//
// RETURNS:
//   - DescribeBlbServerHealthResponse: The return type of the DescribeBlbServerHealth interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbServerHealth(request *DescribeBlbServerHealthRequest) (*DescribeBlbServerHealthResponse, error) {
	result := &DescribeBlbServerHealthResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbServerHealthUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbServers
//
// PARAMS:
//   - request: the arguments to DescribeBlbServers
//
// RETURNS:
//   - DescribeBlbServersResponse: The return type of the DescribeBlbServers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbServers(request *DescribeBlbServersRequest) (*DescribeBlbServersResponse, error) {
	result := &DescribeBlbServersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbServersUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbSslListener
//
// PARAMS:
//   - request: the arguments to DescribeBlbSslListener
//
// RETURNS:
//   - DescribeBlbSslListenerResponse: The return type of the DescribeBlbSslListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbSslListener(request *DescribeBlbSslListenerRequest) (*DescribeBlbSslListenerResponse, error) {
	result := &DescribeBlbSslListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbSslListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbTcpListener
//
// PARAMS:
//   - request: the arguments to DescribeBlbTcpListener
//
// RETURNS:
//   - DescribeBlbTcpListenerResponse: The return type of the DescribeBlbTcpListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbTcpListener(request *DescribeBlbTcpListenerRequest) (*DescribeBlbTcpListenerResponse, error) {
	result := &DescribeBlbTcpListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbTcpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbUdpListener
//
// PARAMS:
//   - request: the arguments to DescribeBlbUdpListener
//
// RETURNS:
//   - DescribeBlbUdpListenerResponse: The return type of the DescribeBlbUdpListener interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbUdpListener(request *DescribeBlbUdpListenerRequest) (*DescribeBlbUdpListenerResponse, error) {
	result := &DescribeBlbUdpListenerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbUdpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeBlbs
//
// PARAMS:
//   - request: the arguments to DescribeBlbs
//
// RETURNS:
//   - DescribeBlbsResponse: The return type of the DescribeBlbs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeBlbs(request *DescribeBlbsRequest) (*DescribeBlbsResponse, error) {
	result := &DescribeBlbsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeBlbsUri(VERSION_V1)).
		WithQueryParamFilter("address", util.StringValue(request.Address)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("blbId", util.StringValue(request.BlbId)).
		WithQueryParamFilter("bccId", util.StringValue(request.BccId)).
		WithQueryParamFilter("exactlyMatch", util.BoolValue(request.ExactlyMatch)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RefundBlb
//
// PARAMS:
//   - request: the arguments to RefundBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RefundBlb(request *RefundBlbRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRefundBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleaseAppBlb
//
// PARAMS:
//   - request: the arguments to ReleaseAppBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseAppBlb(request *ReleaseAppBlbRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseAppBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		Do()
}

// ReleaseBlb
//
// PARAMS:
//   - request: the arguments to ReleaseBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseBlb(request *ReleaseBlbRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ResizeBlb
//
// PARAMS:
//   - request: the arguments to ResizeBlb
//
// RETURNS:
//   - ResizeBlbResponse: The return type of the ResizeBlb interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeBlb(request *ResizeBlbRequest) (*ResizeBlbResponse, error) {
	result := &ResizeBlbResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getResizeBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("action", "RESIZE").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindBlbEnterpriseSecurityGroup
//
// PARAMS:
//   - request: the arguments to UnbindBlbEnterpriseSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindBlbEnterpriseSecurityGroup(request *UnbindBlbEnterpriseSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindBlbEnterpriseSecurityGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UnbindBlbSecurityGroup
//
// PARAMS:
//   - request: the arguments to UnbindBlbSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindBlbSecurityGroup(request *UnbindBlbSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindBlbSecurityGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateAppBlb
//
// PARAMS:
//   - request: the arguments to UpdateAppBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlb(request *UpdateAppBlbRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateAppBlbHttpListener
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbHttpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbHttpListener(request *UpdateAppBlbHttpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbHttpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateAppBlbHttpsListener
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbHttpsListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbHttpsListener(request *UpdateAppBlbHttpsListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbHttpsListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateAppBlbPolicy
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbPolicy(request *UpdateAppBlbPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbPolicyUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("batchupdate", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateAppBlbServerGroup
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbServerGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbServerGroup(request *UpdateAppBlbServerGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbServerGroupUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateAppBlbServerGroupPort
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbServerGroupPort
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbServerGroupPort(request *UpdateAppBlbServerGroupPortRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbServerGroupPortUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateAppBlbServerGroupRs
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbServerGroupRs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbServerGroupRs(request *UpdateAppBlbServerGroupRsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbServerGroupRsUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateAppBlbSslListener
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbSslListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbSslListener(request *UpdateAppBlbSslListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbSslListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateAppBlbTcpListener
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbTcpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbTcpListener(request *UpdateAppBlbTcpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbTcpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateAppBlbUdpListener
//
// PARAMS:
//   - request: the arguments to UpdateAppBlbUdpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAppBlbUdpListener(request *UpdateAppBlbUdpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAppBlbUdpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateBlb
//
// PARAMS:
//   - request: the arguments to UpdateBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlb(request *UpdateBlbRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateBlbAcl
//
// PARAMS:
//   - request: the arguments to UpdateBlbAcl
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbAcl(request *UpdateBlbAclRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbAclUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateBlbHttpListener
//
// PARAMS:
//   - request: the arguments to UpdateBlbHttpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbHttpListener(request *UpdateBlbHttpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbHttpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateBlbHttpsListener
//
// PARAMS:
//   - request: the arguments to UpdateBlbHttpsListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbHttpsListener(request *UpdateBlbHttpsListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbHttpsListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateBlbModifyProtection
//
// PARAMS:
//   - request: the arguments to UpdateBlbModifyProtection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbModifyProtection(request *UpdateBlbModifyProtectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbModifyProtectionUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateBlbServer
//
// PARAMS:
//   - request: the arguments to UpdateBlbServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbServer(request *UpdateBlbServerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbServerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParam("update", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateBlbSslListener
//
// PARAMS:
//   - request: the arguments to UpdateBlbSslListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbSslListener(request *UpdateBlbSslListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbSslListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateBlbTcpListener
//
// PARAMS:
//   - request: the arguments to UpdateBlbTcpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbTcpListener(request *UpdateBlbTcpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbTcpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}

// UpdateBlbUdpListener
//
// PARAMS:
//   - request: the arguments to UpdateBlbUdpListener
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBlbUdpListener(request *UpdateBlbUdpListenerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBlbUdpListenerUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("listenerPort", util.Int32Value(request.ListenerPort)).
		WithBody(request).
		Do()
}
