package blb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "blb." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_BLB = "blb"

	CONSTANT_ACL = "acl"

	CONSTANT_APPBLB = "appblb"

	CONSTANT_H_T_T_P_SLISTENER = "HTTPSlistener"

	CONSTANT_CHARGE = "charge"

	CONSTANT_T_C_PLISTENER = "TCPlistener"

	CONSTANT_LISTENER = "listener"

	CONSTANT_U_D_PLISTENER = "UDPlistener"

	CONSTANT_H_T_T_PLISTENER = "HTTPlistener"

	CONSTANT_POLICYS = "policys"

	CONSTANT_S_S_LLISTENER = "SSLlistener"

	CONSTANT_REFUND = "refund"

	CONSTANT_PRICE = "price"

	CONSTANT_MODIFICATION_PROTECTION = "modification_protection"
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

func getBillingChangeCancelToPostBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_CHARGE
}
func getBillingChangePostToPreBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_CHARGE
}
func getBillingChangePreToPostBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_CHARGE
}
func getBlbInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_PRICE
}
func getCreateAppBlbUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB
}
func getCreateAppBlbHttpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_PLISTENER
}
func getCreateAppBlbHttpsListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_P_SLISTENER
}
func getCreateAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getCreateAppBlbSslListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_S_S_LLISTENER
}
func getCreateAppBlbTcpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_T_C_PLISTENER
}
func getCreateAppBlbUdpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_U_D_PLISTENER
}
func getCreateBlbUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB
}
func getDeleteAppBlbListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_LISTENER
}
func getDeleteAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getDescribeAppBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId
}
func getDescribeAppBlbHttpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_PLISTENER
}
func getDescribeAppBlbHttpsListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_P_SLISTENER
}
func getDescribeAppBlbListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_LISTENER
}
func getDescribeAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getDescribeAppBlbSslListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_S_S_LLISTENER
}
func getDescribeAppBlbTcpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_T_C_PLISTENER
}
func getDescribeAppBlbUdpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_U_D_PLISTENER
}
func getDescribeAppBlbsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB
}
func getDescribeBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId
}
func getDescribeBlbsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB
}
func getRefundBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_REFUND + bce.URI_PREFIX + BlbId
}
func getReleaseAppBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId
}
func getReleaseBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId
}
func getResizeBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId
}
func getUpdateAppBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId
}
func getUpdateAppBlbHttpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_PLISTENER
}
func getUpdateAppBlbHttpsListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_P_SLISTENER
}
func getUpdateAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getUpdateAppBlbSslListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_S_S_LLISTENER
}
func getUpdateAppBlbTcpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_T_C_PLISTENER
}
func getUpdateAppBlbUdpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_U_D_PLISTENER
}
func getUpdateBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId
}
func getUpdateBlbAclUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_ACL + bce.URI_PREFIX + BlbId
}
func getUpdateBlbModifyProtectionUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_MODIFICATION_PROTECTION + bce.URI_PREFIX + BlbId
}
