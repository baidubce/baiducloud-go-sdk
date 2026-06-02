package blb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "blb." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_BLB = "blb"

	CONSTANT_SECURITYGROUP = "securitygroup"

	CONSTANT_ENTERPRISE = "enterprise"

	CONSTANT_APPBLB = "appblb"

	CONSTANT_APPSERVERGROUP = "appservergroup"

	CONSTANT_LBDC = "lbdc"

	CONSTANT_SERVICE = "service"

	CONSTANT_H_T_T_P_SLISTENER = "HTTPSlistener"

	CONSTANT_BACKENDSERVER = "backendserver"

	CONSTANT_IPGROUP = "ipgroup"

	CONSTANT_MEMBER = "member"

	CONSTANT_CHARGE = "charge"

	CONSTANT_T_C_PLISTENER = "TCPlistener"

	CONSTANT_U_D_PLISTENER = "UDPlistener"

	CONSTANT_S_S_LLISTENER = "SSLlistener"

	CONSTANT_LISTENER = "listener"

	CONSTANT_BACKENDPOLICY = "backendpolicy"

	CONSTANT_APPSERVERGROUPPORT = "appservergroupport"

	CONSTANT_REFUND = "refund"

	CONSTANT_BLBRS = "blbrs"

	CONSTANT_H_T_T_PLISTENER = "HTTPlistener"

	CONSTANT_MODIFICATION_PROTECTION = "modification_protection"

	CONSTANT_ACL = "acl"

	CONSTANT_POLICYS = "policys"

	CONSTANT_PRICE = "price"

	CONSTANT_BLBRSUNMOUNT = "blbrsunmount"

	CONSTANT_BLBRSMOUNT = "blbrsmount"
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

func getAddAppBlbServerGroupRsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BLBRS
}
func getAddBlbServerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BACKENDSERVER
}
func getAddServiceAuthUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
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
func getBindBlbEnterpriseSecurityGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITYGROUP
}
func getBindBlbSecurityGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_SECURITYGROUP
}
func getBindInstanceToServiceUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
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
func getCreateAppBlbIpGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP
}
func getCreateAppBlbIpGroupMemberUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_MEMBER
}
func getCreateAppBlbIpGroupProtocolUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_BACKENDPOLICY
}
func getCreateAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getCreateAppBlbServerGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUP
}
func getCreateAppBlbServerGroupPortUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUPPORT
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
func getCreateBlbHttpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_PLISTENER
}
func getCreateBlbHttpsListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_P_SLISTENER
}
func getCreateBlbSslListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_S_S_LLISTENER
}
func getCreateBlbTcpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_T_C_PLISTENER
}
func getCreateBlbUdpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_U_D_PLISTENER
}
func getCreateLbdcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC
}
func getCreateServiceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE
}
func getDeleteAppBlbIpGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP
}
func getDeleteAppBlbIpGroupMemberUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_MEMBER
}
func getDeleteAppBlbIpGroupProtocolUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_BACKENDPOLICY
}
func getDeleteAppBlbListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_LISTENER
}
func getDeleteAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getDeleteAppBlbServerGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUP
}
func getDeleteAppBlbServerGroupPortUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUPPORT
}
func getDeleteAppBlbServerGroupRsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BLBRS
}
func getDeleteBlbListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_LISTENER
}
func getDeleteBlbServerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BACKENDSERVER
}
func getDeleteServiceUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
}
func getDeleteServiceAuthUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
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
func getDescribeAppBlbIpGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP
}
func getDescribeAppBlbIpGroupMemberUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_MEMBER
}
func getDescribeAppBlbListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_LISTENER
}
func getDescribeAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getDescribeAppBlbServerGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUP
}
func getDescribeAppBlbServerGroupMountRsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BLBRSMOUNT
}
func getDescribeAppBlbServerGroupRsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BLBRS
}
func getDescribeAppBlbServerGroupUnmountRsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BLBRSUNMOUNT
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
func getDescribeBlbEnterpriseSecurityGroupsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITYGROUP
}
func getDescribeBlbHttpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_PLISTENER
}
func getDescribeBlbHttpsListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_P_SLISTENER
}
func getDescribeBlbListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_LISTENER
}
func getDescribeBlbSecurityGroupsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_SECURITYGROUP
}
func getDescribeBlbServerHealthUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BACKENDSERVER
}
func getDescribeBlbServersUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BACKENDSERVER
}
func getDescribeBlbSslListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_S_S_LLISTENER
}
func getDescribeBlbTcpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_T_C_PLISTENER
}
func getDescribeBlbUdpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_U_D_PLISTENER
}
func getDescribeBlbsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB
}
func getDescribeLbdcUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC + bce.URI_PREFIX + Id
}
func getDescribeLbdcBlbUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_BLB
}
func getDescribeLbdcsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC
}
func getDescribeServiceUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
}
func getDescribeServicesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE
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
func getRenewLbdcUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC + bce.URI_PREFIX + Id
}
func getResizeBlbUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId
}
func getUnbindBlbEnterpriseSecurityGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITYGROUP
}
func getUnbindBlbSecurityGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_SECURITYGROUP
}
func getUnbindInstanceFromServiceUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
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
func getUpdateAppBlbIpGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP
}
func getUpdateAppBlbIpGroupMemberUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_MEMBER
}
func getUpdateAppBlbIpGroupProtocolUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_IPGROUP + bce.URI_PREFIX + CONSTANT_BACKENDPOLICY
}
func getUpdateAppBlbPolicyUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_POLICYS
}
func getUpdateAppBlbServerGroupUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUP
}
func getUpdateAppBlbServerGroupPortUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_APPSERVERGROUPPORT
}
func getUpdateAppBlbServerGroupRsUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APPBLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BLBRS
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
func getUpdateBlbHttpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_PLISTENER
}
func getUpdateBlbHttpsListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_H_T_T_P_SLISTENER
}
func getUpdateBlbModifyProtectionUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + CONSTANT_MODIFICATION_PROTECTION + bce.URI_PREFIX + BlbId
}
func getUpdateBlbServerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_BACKENDSERVER
}
func getUpdateBlbSslListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_S_S_LLISTENER
}
func getUpdateBlbTcpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_T_C_PLISTENER
}
func getUpdateBlbUdpListenerUri(version string, BlbId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLB + bce.URI_PREFIX + BlbId + bce.URI_PREFIX + CONSTANT_U_D_PLISTENER
}
func getUpdateLbdcUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC + bce.URI_PREFIX + Id
}
func getUpdateServiceUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
}
func getUpdateServiceAuthUri(version string, Service string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SERVICE + bce.URI_PREFIX + Service
}
func getUpgradeLbdcUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LBDC + bce.URI_PREFIX + Id
}
