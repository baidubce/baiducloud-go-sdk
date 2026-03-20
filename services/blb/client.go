package blb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "blb." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_BLB = "blb"

	CONSTANT_CHARGE = "charge"

	CONSTANT_PRICE = "price"

	CONSTANT_REFUND = "refund"
)

// Client of blb service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getBillingChangeCancelToPostBlbUri(version string, blbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + blbId + bce.URI_PREFIX + CONSTANT_CHARGE
}
func getBillingChangePostToPreBlbUri(version string, blbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + blbId + bce.URI_PREFIX + CONSTANT_CHARGE
}
func getBillingChangePreToPostBlbUri(version string, blbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + blbId + bce.URI_PREFIX + CONSTANT_CHARGE
}
func getBlbInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_PRICE
}
func getRefundBlbUri(version string, blbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_REFUND + bce.URI_PREFIX + blbId
}
func getResizeBlbUri(version string, blbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + blbId
}
