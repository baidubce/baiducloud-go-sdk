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
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBillingChangeCancelToPostBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("action", "CANCEL_TO_POSTPAY").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
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
	return result, err
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
	return result, err
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
	return result, err
}

// RefundBlb
//
// PARAMS:
//   - request: the arguments to RefundBlb
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RefundBlb(request *RefundBlbRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRefundBlbUri(VERSION_V1, util.StringValue(request.BlbId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
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
	return result, err
}
