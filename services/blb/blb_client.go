package blb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

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
