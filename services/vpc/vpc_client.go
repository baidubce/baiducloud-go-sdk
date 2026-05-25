package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"strings"
)

const (
	VERSION_V1 = "v1"

	VERSION_V2 = "v2"
)

// AcceptPeerToPeerConnectionApplications
//
// PARAMS:
//   - request: the arguments to AcceptPeerToPeerConnectionApplications
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AcceptPeerToPeerConnectionApplications(request *AcceptPeerToPeerConnectionApplicationsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAcceptPeerToPeerConnectionApplicationsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("accept", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ActiveStandbySwitchover
//
// PARAMS:
//   - request: the arguments to ActiveStandbySwitchover
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ActiveStandbySwitchover(request *ActiveStandbySwitchoverRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getActiveStandbySwitchoverUri(VERSION_V1, util.StringValue(request.RouteRuleId))).
		WithQueryParam("switchRouteHA", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// AddAclRule
//
// PARAMS:
//   - request: the arguments to AddAclRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddAclRule(request *AddAclRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddAclRuleUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddElasticNetworkCardAuxiliaryIp
//
// PARAMS:
//   - request: the arguments to AddElasticNetworkCardAuxiliaryIp
//
// RETURNS:
//   - AddElasticNetworkCardAuxiliaryIpResponse: The return type of the AddElasticNetworkCardAuxiliaryIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddElasticNetworkCardAuxiliaryIp(request *AddElasticNetworkCardAuxiliaryIpRequest) (*AddElasticNetworkCardAuxiliaryIpResponse, error) {
	result := &AddElasticNetworkCardAuxiliaryIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddElasticNetworkCardAuxiliaryIpUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddIpAddressGroupToIpAddressFamily
//
// PARAMS:
//   - request: the arguments to AddIpAddressGroupToIpAddressFamily
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddIpAddressGroupToIpAddressFamily(request *AddIpAddressGroupToIpAddressFamilyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddIpAddressGroupToIpAddressFamilyUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddIpAddressesToTheIpAddressGroup
//
// PARAMS:
//   - request: the arguments to AddIpAddressesToTheIpAddressGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddIpAddressesToTheIpAddressGroup(request *AddIpAddressesToTheIpAddressGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddIpAddressesToTheIpAddressGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddIpv6OnlyOutboundAndNoInboundPolicy
//
// PARAMS:
//   - request: the arguments to AddIpv6OnlyOutboundAndNoInboundPolicy
//
// RETURNS:
//   - AddIpv6OnlyOutboundAndNoInboundPolicyResponse: The return type of the AddIpv6OnlyOutboundAndNoInboundPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddIpv6OnlyOutboundAndNoInboundPolicy(request *AddIpv6OnlyOutboundAndNoInboundPolicyRequest) (*AddIpv6OnlyOutboundAndNoInboundPolicyResponse, error) {
	result := &AddIpv6OnlyOutboundAndNoInboundPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddIpv6OnlyOutboundAndNoInboundPolicyUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AuthorizeRegularSecurityGroupRulesV2
//
// PARAMS:
//   - request: the arguments to AuthorizeRegularSecurityGroupRulesV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuthorizeRegularSecurityGroupRulesV2(request *AuthorizeRegularSecurityGroupRulesV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAuthorizeRegularSecurityGroupRulesV2Uri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		WithQueryParam("authorizeRule", "").
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AuthorizedEnterpriseSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to AuthorizedEnterpriseSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuthorizedEnterpriseSecurityGroupRules(request *AuthorizedEnterpriseSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAuthorizedEnterpriseSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.EnterpriseSecurityGroupId))).
		WithQueryParam("authorizeRule", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BatchCreateSslVpnUsers
//
// PARAMS:
//   - request: the arguments to BatchCreateSslVpnUsers
//
// RETURNS:
//   - BatchCreateSslVpnUsersResponse: The return type of the BatchCreateSslVpnUsers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchCreateSslVpnUsers(request *BatchCreateSslVpnUsersRequest) (*BatchCreateSslVpnUsersResponse, error) {
	result := &BatchCreateSslVpnUsersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchCreateSslVpnUsersUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchDeleteElasticNetworkCardIntranetIp
//
// PARAMS:
//   - request: the arguments to BatchDeleteElasticNetworkCardIntranetIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteElasticNetworkCardIntranetIp(request *BatchDeleteElasticNetworkCardIntranetIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchDeleteElasticNetworkCardIntranetIpUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BatchIncreaseElasticNetworkCardIntranetIp
//
// PARAMS:
//   - request: the arguments to BatchIncreaseElasticNetworkCardIntranetIp
//
// RETURNS:
//   - BatchIncreaseElasticNetworkCardIntranetIpResponse: The return type of the BatchIncreaseElasticNetworkCardIntranetIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchIncreaseElasticNetworkCardIntranetIp(request *BatchIncreaseElasticNetworkCardIntranetIpRequest) (*BatchIncreaseElasticNetworkCardIntranetIpResponse, error) {
	result := &BatchIncreaseElasticNetworkCardIntranetIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchIncreaseElasticNetworkCardIntranetIpUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindEip
//
// PARAMS:
//   - request: the arguments to BindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindEip(request *BindEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindEipUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindPhysicalDedicatedLine
//
// PARAMS:
//   - request: the arguments to BindPhysicalDedicatedLine
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindPhysicalDedicatedLine(request *BindPhysicalDedicatedLineRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindPhysicalDedicatedLineUri(VERSION_V1, util.StringValue(request.EtGatewayId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ClosePeerToPeerConnectionToSynchronizeDns
//
// PARAMS:
//   - request: the arguments to ClosePeerToPeerConnectionToSynchronizeDns
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ClosePeerToPeerConnectionToSynchronizeDns(request *ClosePeerToPeerConnectionToSynchronizeDnsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClosePeerToPeerConnectionToSynchronizeDnsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("close", "").
		WithQueryParamFilter("role", util.StringValue(request.Role)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// CloseVpcRelay
//
// PARAMS:
//   - request: the arguments to CloseVpcRelay
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CloseVpcRelay(request *CloseVpcRelayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCloseVpcRelayUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// CreateAHighlyAvailableVirtualIp
//
// PARAMS:
//   - request: the arguments to CreateAHighlyAvailableVirtualIp
//
// RETURNS:
//   - CreateAHighlyAvailableVirtualIpResponse: The return type of the CreateAHighlyAvailableVirtualIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAHighlyAvailableVirtualIp(request *CreateAHighlyAvailableVirtualIpRequest) (*CreateAHighlyAvailableVirtualIpResponse, error) {
	result := &CreateAHighlyAvailableVirtualIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAHighlyAvailableVirtualIpUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAPeerToPeerConnection
//
// PARAMS:
//   - request: the arguments to CreateAPeerToPeerConnection
//
// RETURNS:
//   - CreateAPeerToPeerConnectionResponse: The return type of the CreateAPeerToPeerConnection interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAPeerToPeerConnection(request *CreateAPeerToPeerConnectionRequest) (*CreateAPeerToPeerConnectionResponse, error) {
	result := &CreateAPeerToPeerConnectionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAPeerToPeerConnectionUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateARegularSecurityGroupV2
//
// PARAMS:
//   - request: the arguments to CreateARegularSecurityGroupV2
//
// RETURNS:
//   - CreateARegularSecurityGroupV2Response: The return type of the CreateARegularSecurityGroupV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateARegularSecurityGroupV2(request *CreateARegularSecurityGroupV2Request) (*CreateARegularSecurityGroupV2Response, error) {
	result := &CreateARegularSecurityGroupV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateARegularSecurityGroupV2Uri(VERSION_V2)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAnIpv6Gateway
//
// PARAMS:
//   - request: the arguments to CreateAnIpv6Gateway
//
// RETURNS:
//   - CreateAnIpv6GatewayResponse: The return type of the CreateAnIpv6Gateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAnIpv6Gateway(request *CreateAnIpv6GatewayRequest) (*CreateAnIpv6GatewayResponse, error) {
	result := &CreateAnIpv6GatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAnIpv6GatewayUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDedicatedGateway
//
// PARAMS:
//   - request: the arguments to CreateDedicatedGateway
//
// RETURNS:
//   - CreateDedicatedGatewayResponse: The return type of the CreateDedicatedGateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedGateway(request *CreateDedicatedGatewayRequest) (*CreateDedicatedGatewayResponse, error) {
	result := &CreateDedicatedGatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDedicatedGatewayUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDedicatedGatewayHealthCheck
//
// PARAMS:
//   - request: the arguments to CreateDedicatedGatewayHealthCheck
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedGatewayHealthCheck(request *CreateDedicatedGatewayHealthCheckRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDedicatedGatewayHealthCheckUri(VERSION_V1, util.StringValue(request.EtGatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateElasticNetworkCard
//
// PARAMS:
//   - request: the arguments to CreateElasticNetworkCard
//
// RETURNS:
//   - CreateElasticNetworkCardResponse: The return type of the CreateElasticNetworkCard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateElasticNetworkCard(request *CreateElasticNetworkCardRequest) (*CreateElasticNetworkCardResponse, error) {
	result := &CreateElasticNetworkCardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateElasticNetworkCardUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEnterpriseSecurityGroup
//
// PARAMS:
//   - request: the arguments to CreateEnterpriseSecurityGroup
//
// RETURNS:
//   - CreateEnterpriseSecurityGroupResponse: The return type of the CreateEnterpriseSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEnterpriseSecurityGroup(request *CreateEnterpriseSecurityGroupRequest) (*CreateEnterpriseSecurityGroupResponse, error) {
	result := &CreateEnterpriseSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEnterpriseSecurityGroupUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateGatewayLimitRules
//
// PARAMS:
//   - request: the arguments to CreateGatewayLimitRules
//
// RETURNS:
//   - CreateGatewayLimitRulesResponse: The return type of the CreateGatewayLimitRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateGatewayLimitRules(request *CreateGatewayLimitRulesRequest) (*CreateGatewayLimitRulesResponse, error) {
	result := &CreateGatewayLimitRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateGatewayLimitRulesUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIpAddressFamily
//
// PARAMS:
//   - request: the arguments to CreateIpAddressFamily
//
// RETURNS:
//   - CreateIpAddressFamilyResponse: The return type of the CreateIpAddressFamily interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpAddressFamily(request *CreateIpAddressFamilyRequest) (*CreateIpAddressFamilyResponse, error) {
	result := &CreateIpAddressFamilyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpAddressFamilyUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIpAddressGroup
//
// PARAMS:
//   - request: the arguments to CreateIpAddressGroup
//
// RETURNS:
//   - CreateIpAddressGroupResponse: The return type of the CreateIpAddressGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpAddressGroup(request *CreateIpAddressGroupRequest) (*CreateIpAddressGroupResponse, error) {
	result := &CreateIpAddressGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpAddressGroupUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIpReserved
//
// PARAMS:
//   - request: the arguments to CreateIpReserved
//
// RETURNS:
//   - CreateIpReservedResponse: The return type of the CreateIpReserved interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpReserved(request *CreateIpReservedRequest) (*CreateIpReservedResponse, error) {
	result := &CreateIpReservedResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpReservedUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIpv6GatewaySpeedLimitPolicy
//
// PARAMS:
//   - request: the arguments to CreateIpv6GatewaySpeedLimitPolicy
//
// RETURNS:
//   - CreateIpv6GatewaySpeedLimitPolicyResponse: The return type of the CreateIpv6GatewaySpeedLimitPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpv6GatewaySpeedLimitPolicy(request *CreateIpv6GatewaySpeedLimitPolicyRequest) (*CreateIpv6GatewaySpeedLimitPolicyResponse, error) {
	result := &CreateIpv6GatewaySpeedLimitPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpv6GatewaySpeedLimitPolicyUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateNetworkDetection
//
// PARAMS:
//   - request: the arguments to CreateNetworkDetection
//
// RETURNS:
//   - CreateNetworkDetectionResponse: The return type of the CreateNetworkDetection interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateNetworkDetection(request *CreateNetworkDetectionRequest) (*CreateNetworkDetectionResponse, error) {
	result := &CreateNetworkDetectionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateNetworkDetectionUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRoutingRules
//
// PARAMS:
//   - request: the arguments to CreateRoutingRules
//
// RETURNS:
//   - CreateRoutingRulesResponse: The return type of the CreateRoutingRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateRoutingRules(request *CreateRoutingRulesRequest) (*CreateRoutingRulesResponse, error) {
	result := &CreateRoutingRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRoutingRulesUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSslVpnServer
//
// PARAMS:
//   - request: the arguments to CreateSslVpnServer
//
// RETURNS:
//   - CreateSslVpnServerResponse: The return type of the CreateSslVpnServer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSslVpnServer(request *CreateSslVpnServerRequest) (*CreateSslVpnServerResponse, error) {
	result := &CreateSslVpnServerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSslVpnServerUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSubnet
//
// PARAMS:
//   - request: the arguments to CreateSubnet
//
// RETURNS:
//   - CreateSubnetResponse: The return type of the CreateSubnet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (*CreateSubnetResponse, error) {
	result := &CreateSubnetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSubnetUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUserGateway
//
// PARAMS:
//   - request: the arguments to CreateUserGateway
//
// RETURNS:
//   - CreateUserGatewayResponse: The return type of the CreateUserGateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateUserGateway(request *CreateUserGatewayRequest) (*CreateUserGatewayResponse, error) {
	result := &CreateUserGatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateUserGatewayUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateVpc
//
// PARAMS:
//   - request: the arguments to CreateVpc
//
// RETURNS:
//   - CreateVpcResponse: The return type of the CreateVpc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVpc(request *CreateVpcRequest) (*CreateVpcResponse, error) {
	result := &CreateVpcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateVpcUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateVpn
//
// PARAMS:
//   - request: the arguments to CreateVpn
//
// RETURNS:
//   - CreateVpnResponse: The return type of the CreateVpn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVpn(request *CreateVpnRequest) (*CreateVpnResponse, error) {
	result := &CreateVpnResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateVpnUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateVpnTunnel
//
// PARAMS:
//   - request: the arguments to CreateVpnTunnel
//
// RETURNS:
//   - CreateVpnTunnelResponse: The return type of the CreateVpnTunnel interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVpnTunnel(request *CreateVpnTunnelRequest) (*CreateVpnTunnelResponse, error) {
	result := &CreateVpnTunnelResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateVpnTunnelUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAclRule
//
// PARAMS:
//   - request: the arguments to DeleteAclRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAclRule(request *DeleteAclRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAclRuleUri(VERSION_V1, util.StringValue(request.AclRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteElasticNetworkCardAuxiliaryIp
//
// PARAMS:
//   - request: the arguments to DeleteElasticNetworkCardAuxiliaryIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteElasticNetworkCardAuxiliaryIp(request *DeleteElasticNetworkCardAuxiliaryIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteElasticNetworkCardAuxiliaryIpUri(VERSION_V1, util.StringValue(request.EniId), util.StringValue(request.PrivateIpAddress))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteEnterpriseSecurityGroup
//
// PARAMS:
//   - request: the arguments to DeleteEnterpriseSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteEnterpriseSecurityGroup(request *DeleteEnterpriseSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteEnterpriseSecurityGroupUri(VERSION_V1, util.StringValue(request.EnterpriseSecurityGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteEnterpriseSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to DeleteEnterpriseSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteEnterpriseSecurityGroupRules(request *DeleteEnterpriseSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteEnterpriseSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.EnterpriseSecurityGroupRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteGatewayLimitRule
//
// PARAMS:
//   - request: the arguments to DeleteGatewayLimitRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteGatewayLimitRule(request *DeleteGatewayLimitRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteGatewayLimitRuleUri(VERSION_V1, util.StringValue(request.GlrId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteHighlyAvailableVirtualIp
//
// PARAMS:
//   - request: the arguments to DeleteHighlyAvailableVirtualIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteHighlyAvailableVirtualIp(request *DeleteHighlyAvailableVirtualIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteHighlyAvailableVirtualIpUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpAddressFamily
//
// PARAMS:
//   - request: the arguments to DeleteIpAddressFamily
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpAddressFamily(request *DeleteIpAddressFamilyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpAddressFamilyUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpAddressFromIpAddressGroup
//
// PARAMS:
//   - request: the arguments to DeleteIpAddressFromIpAddressGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpAddressFromIpAddressGroup(request *DeleteIpAddressFromIpAddressGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteIpAddressFromIpAddressGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteIpAddressGroup
//
// PARAMS:
//   - request: the arguments to DeleteIpAddressGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpAddressGroup(request *DeleteIpAddressGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpAddressGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpReserve
//
// PARAMS:
//   - request: the arguments to DeleteIpReserve
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpReserve(request *DeleteIpReserveRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpReserveUri(VERSION_V1, util.StringValue(request.IpReserveId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpv6Gateway
//
// PARAMS:
//   - request: the arguments to DeleteIpv6Gateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpv6Gateway(request *DeleteIpv6GatewayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpv6GatewayUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpv6GatewaySpeedLimitPolicy
//
// PARAMS:
//   - request: the arguments to DeleteIpv6GatewaySpeedLimitPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpv6GatewaySpeedLimitPolicy(request *DeleteIpv6GatewaySpeedLimitPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpv6GatewaySpeedLimitPolicyUri(VERSION_V1, util.StringValue(request.GatewayId), util.StringValue(request.RateLimitRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpv6OnlyAccessPolicy
//
// PARAMS:
//   - request: the arguments to DeleteIpv6OnlyAccessPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpv6OnlyAccessPolicy(request *DeleteIpv6OnlyAccessPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpv6OnlyAccessPolicyUri(VERSION_V1, util.StringValue(request.GatewayId), util.StringValue(request.EgressOnlyRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteNetworkDetection
//
// PARAMS:
//   - request: the arguments to DeleteNetworkDetection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteNetworkDetection(request *DeleteNetworkDetectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteNetworkDetectionUri(VERSION_V1, util.StringValue(request.ProbeId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteRegularSecurityGroupRulesV2
//
// PARAMS:
//   - request: the arguments to DeleteRegularSecurityGroupRulesV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRegularSecurityGroupRulesV2(request *DeleteRegularSecurityGroupRulesV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRegularSecurityGroupRulesV2Uri(VERSION_V2, util.StringValue(request.SecurityGroupRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		Do()
}

// DeleteRegularSecurityGroupV2
//
// PARAMS:
//   - request: the arguments to DeleteRegularSecurityGroupV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRegularSecurityGroupV2(request *DeleteRegularSecurityGroupV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRegularSecurityGroupV2Uri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteRoutingRules
//
// PARAMS:
//   - request: the arguments to DeleteRoutingRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRoutingRules(request *DeleteRoutingRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRoutingRulesUri(VERSION_V1, util.StringValue(request.RouteRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteSslVpnServer
//
// PARAMS:
//   - request: the arguments to DeleteSslVpnServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSslVpnServer(request *DeleteSslVpnServerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSslVpnServerUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.SslVpnServerId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteSslVpnUser
//
// PARAMS:
//   - request: the arguments to DeleteSslVpnUser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSslVpnUser(request *DeleteSslVpnUserRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSslVpnUserUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.UserId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteSubnet
//
// PARAMS:
//   - request: the arguments to DeleteSubnet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteUserGateway
//
// PARAMS:
//   - request: the arguments to DeleteUserGateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteUserGateway(request *DeleteUserGatewayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteUserGatewayUri(VERSION_V1, util.StringValue(request.CgwId))).
		Do()
}

// DeleteVpc
//
// PARAMS:
//   - request: the arguments to DeleteVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteVpc(request *DeleteVpcRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteVpnTunnel
//
// PARAMS:
//   - request: the arguments to DeleteVpnTunnel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteVpnTunnel(request *DeleteVpnTunnelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteVpnTunnelUri(VERSION_V1, util.StringValue(request.VpnConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ElasticNetworkCardBindingEip
//
// PARAMS:
//   - request: the arguments to ElasticNetworkCardBindingEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ElasticNetworkCardBindingEip(request *ElasticNetworkCardBindingEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getElasticNetworkCardBindingEipUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ElasticNetworkCardMountedCloudProductInstance
//
// PARAMS:
//   - request: the arguments to ElasticNetworkCardMountedCloudProductInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ElasticNetworkCardMountedCloudProductInstance(request *ElasticNetworkCardMountedCloudProductInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getElasticNetworkCardMountedCloudProductInstanceUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("attach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ElasticNetworkCardUnbindingEip
//
// PARAMS:
//   - request: the arguments to ElasticNetworkCardUnbindingEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ElasticNetworkCardUnbindingEip(request *ElasticNetworkCardUnbindingEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getElasticNetworkCardUnbindingEipUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("unBind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ElasticNetworkCardUninstallationCloudProductInstance
//
// PARAMS:
//   - request: the arguments to ElasticNetworkCardUninstallationCloudProductInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ElasticNetworkCardUninstallationCloudProductInstance(request *ElasticNetworkCardUninstallationCloudProductInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getElasticNetworkCardUninstallationCloudProductInstanceUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("detach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ElasticNetworkCardUpdateEnterpriseSecurityGroup
//
// PARAMS:
//   - request: the arguments to ElasticNetworkCardUpdateEnterpriseSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ElasticNetworkCardUpdateEnterpriseSecurityGroup(request *ElasticNetworkCardUpdateEnterpriseSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getElasticNetworkCardUpdateEnterpriseSecurityGroupUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("bindEsg", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ElasticNetworkCardUpdatesRegularSecurityGroup
//
// PARAMS:
//   - request: the arguments to ElasticNetworkCardUpdatesRegularSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ElasticNetworkCardUpdatesRegularSecurityGroup(request *ElasticNetworkCardUpdatesRegularSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getElasticNetworkCardUpdatesRegularSecurityGroupUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("bindSg", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// EnablePeerToPeerConnectionToSynchronizeDns
//
// PARAMS:
//   - request: the arguments to EnablePeerToPeerConnectionToSynchronizeDns
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnablePeerToPeerConnectionToSynchronizeDns(request *EnablePeerToPeerConnectionToSynchronizeDnsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnablePeerToPeerConnectionToSynchronizeDnsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("open", "").
		WithQueryParamFilter("role", util.StringValue(request.Role)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// GetVpcResourceIpInfo
//
// PARAMS:
//   - request: the arguments to GetVpcResourceIpInfo
//
// RETURNS:
//   - GetVpcResourceIpInfoResponse: The return type of the GetVpcResourceIpInfo interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetVpcResourceIpInfo(request *GetVpcResourceIpInfoRequest) (*GetVpcResourceIpInfoResponse, error) {
	result := &GetVpcResourceIpInfoResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetVpcResourceIpInfoUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("subnetId", util.StringValue(request.SubnetId)).
		WithQueryParamFilter("resourceType", util.StringValue(request.ResourceType)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HighAvailabilityVirtualIpUnbindingEip
//
// PARAMS:
//   - request: the arguments to HighAvailabilityVirtualIpUnbindingEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HighAvailabilityVirtualIpUnbindingEip(request *HighAvailabilityVirtualIpUnbindingEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHighAvailabilityVirtualIpUnbindingEipUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("unbindPublicIp", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// HighAvailabilityVirtualIpUnbindingInstance
//
// PARAMS:
//   - request: the arguments to HighAvailabilityVirtualIpUnbindingInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HighAvailabilityVirtualIpUnbindingInstance(request *HighAvailabilityVirtualIpUnbindingInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHighAvailabilityVirtualIpUnbindingInstanceUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("detach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// HighlyAvailableVirtualIpBindingEip
//
// PARAMS:
//   - request: the arguments to HighlyAvailableVirtualIpBindingEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HighlyAvailableVirtualIpBindingEip(request *HighlyAvailableVirtualIpBindingEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHighlyAvailableVirtualIpBindingEipUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("bindPublicIp", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// HighlyAvailableVirtualIpBindingInstance
//
// PARAMS:
//   - request: the arguments to HighlyAvailableVirtualIpBindingInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HighlyAvailableVirtualIpBindingInstance(request *HighlyAvailableVirtualIpBindingInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHighlyAvailableVirtualIpBindingInstanceUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("attach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// Ipv6GatewayBandwidthUpgradeAndDowngrade
//
// PARAMS:
//   - request: the arguments to Ipv6GatewayBandwidthUpgradeAndDowngrade
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) Ipv6GatewayBandwidthUpgradeAndDowngrade(request *Ipv6GatewayBandwidthUpgradeAndDowngradeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getIpv6GatewayBandwidthUpgradeAndDowngradeUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ListIpReserve
//
// PARAMS:
//   - request: the arguments to ListIpReserve
//
// RETURNS:
//   - ListIpReserveResponse: The return type of the ListIpReserve interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListIpReserve(request *ListIpReserveRequest) (*ListIpReserveResponse, error) {
	result := &ListIpReserveResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListIpReserveUri(VERSION_V1)).
		WithQueryParamFilter("subnetId", util.StringValue(request.SubnetId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyGatewayLimitRules
//
// PARAMS:
//   - request: the arguments to ModifyGatewayLimitRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyGatewayLimitRules(request *ModifyGatewayLimitRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyGatewayLimitRulesUri(VERSION_V1, util.StringValue(request.GlrId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// OpenVpcRelay
//
// PARAMS:
//   - request: the arguments to OpenVpcRelay
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) OpenVpcRelay(request *OpenVpcRelayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getOpenVpcRelayUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// PeerToPeerConnectionBandwidthUpgradeAndDowngrade
//
// PARAMS:
//   - request: the arguments to PeerToPeerConnectionBandwidthUpgradeAndDowngrade
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PeerToPeerConnectionBandwidthUpgradeAndDowngrade(request *PeerToPeerConnectionBandwidthUpgradeAndDowngradeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPeerToPeerConnectionBandwidthUpgradeAndDowngradeUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// PeerToPeerConnectionRenewal
//
// PARAMS:
//   - request: the arguments to PeerToPeerConnectionRenewal
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PeerToPeerConnectionRenewal(request *PeerToPeerConnectionRenewalRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPeerToPeerConnectionRenewalUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// PrepaidPeerToPeerConnectionUnsubscribe
//
// PARAMS:
//   - request: the arguments to PrepaidPeerToPeerConnectionUnsubscribe
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PrepaidPeerToPeerConnectionUnsubscribe(request *PrepaidPeerToPeerConnectionUnsubscribeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPrepaidPeerToPeerConnectionUnsubscribeUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("refund", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// QueryAcl
//
// PARAMS:
//   - request: the arguments to QueryAcl
//
// RETURNS:
//   - QueryAclResponse: The return type of the QueryAcl interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAcl(request *QueryAclRequest) (*QueryAclResponse, error) {
	result := &QueryAclResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAclUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAclRules
//
// PARAMS:
//   - request: the arguments to QueryAclRules
//
// RETURNS:
//   - QueryAclRulesResponse: The return type of the QueryAclRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAclRules(request *QueryAclRulesRequest) (*QueryAclRulesResponse, error) {
	result := &QueryAclRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAclRulesUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("subnetId", util.StringValue(request.SubnetId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryIpAddressFamilyList
//
// PARAMS:
//   - request: the arguments to QueryIpAddressFamilyList
//
// RETURNS:
//   - QueryIpAddressFamilyListResponse: The return type of the QueryIpAddressFamilyList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpAddressFamilyList(request *QueryIpAddressFamilyListRequest) (*QueryIpAddressFamilyListResponse, error) {
	result := &QueryIpAddressFamilyListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpAddressFamilyListUri(VERSION_V1)).
		WithQueryParamFilter("ipVersion", util.StringValue(request.IpVersion)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryIpv6Gateway
//
// PARAMS:
//   - request: the arguments to QueryIpv6Gateway
//
// RETURNS:
//   - QueryIpv6GatewayResponse: The return type of the QueryIpv6Gateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpv6Gateway(request *QueryIpv6GatewayRequest) (*QueryIpv6GatewayResponse, error) {
	result := &QueryIpv6GatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpv6GatewayUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryNetworkDetectionDetails
//
// PARAMS:
//   - request: the arguments to QueryNetworkDetectionDetails
//
// RETURNS:
//   - QueryNetworkDetectionDetailsResponse: The return type of the QueryNetworkDetectionDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryNetworkDetectionDetails(request *QueryNetworkDetectionDetailsRequest) (*QueryNetworkDetectionDetailsResponse, error) {
	result := &QueryNetworkDetectionDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryNetworkDetectionDetailsUri(VERSION_V1, util.StringValue(request.ProbeId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryNetworkDetectionList
//
// PARAMS:
//   - request: the arguments to QueryNetworkDetectionList
//
// RETURNS:
//   - QueryNetworkDetectionListResponse: The return type of the QueryNetworkDetectionList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryNetworkDetectionList(request *QueryNetworkDetectionListRequest) (*QueryNetworkDetectionListResponse, error) {
	result := &QueryNetworkDetectionListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryNetworkDetectionListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRoutingRules
//
// PARAMS:
//   - request: the arguments to QueryRoutingRules
//
// RETURNS:
//   - QueryRoutingRulesResponse: The return type of the QueryRoutingRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRoutingRules(request *QueryRoutingRulesRequest) (*QueryRoutingRulesResponse, error) {
	result := &QueryRoutingRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRoutingRulesUri(VERSION_V1)).
		WithQueryParamFilter("routeTableId", util.StringValue(request.RouteTableId)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRoutingTable
//
// PARAMS:
//   - request: the arguments to QueryRoutingTable
//
// RETURNS:
//   - QueryRoutingTableResponse: The return type of the QueryRoutingTable interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRoutingTable(request *QueryRoutingTableRequest) (*QueryRoutingTableResponse, error) {
	result := &QueryRoutingTableResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRoutingTableUri(VERSION_V1)).
		WithQueryParamFilter("routeTableId", util.StringValue(request.RouteTableId)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySpecifiedSubnet
//
// PARAMS:
//   - request: the arguments to QuerySpecifiedSubnet
//
// RETURNS:
//   - QuerySpecifiedSubnetResponse: The return type of the QuerySpecifiedSubnet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySpecifiedSubnet(request *QuerySpecifiedSubnetRequest) (*QuerySpecifiedSubnetResponse, error) {
	result := &QuerySpecifiedSubnetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySpecifiedSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySpecifiedVpc
//
// PARAMS:
//   - request: the arguments to QuerySpecifiedVpc
//
// RETURNS:
//   - QuerySpecifiedVpcResponse: The return type of the QuerySpecifiedVpc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySpecifiedVpc(request *QuerySpecifiedVpcRequest) (*QuerySpecifiedVpcResponse, error) {
	result := &QuerySpecifiedVpcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySpecifiedVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySslVpnServer
//
// PARAMS:
//   - request: the arguments to QuerySslVpnServer
//
// RETURNS:
//   - QuerySslVpnServerResponse: The return type of the QuerySslVpnServer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySslVpnServer(request *QuerySslVpnServerRequest) (*QuerySslVpnServerResponse, error) {
	result := &QuerySslVpnServerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySslVpnServerUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySslVpnUsers
//
// PARAMS:
//   - request: the arguments to QuerySslVpnUsers
//
// RETURNS:
//   - QuerySslVpnUsersResponse: The return type of the QuerySslVpnUsers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySslVpnUsers(request *QuerySslVpnUsersRequest) (*QuerySslVpnUsersResponse, error) {
	result := &QuerySslVpnUsersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySslVpnUsersUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("userName", util.StringValue(request.UserName)).
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
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("subnetType", util.StringValue(request.SubnetType)).
		WithQueryParamFilter("subnetIds", util.StringValue(request.SubnetIds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheDetailsOfTheDedicatedGateway
//
// PARAMS:
//   - request: the arguments to QueryTheDetailsOfTheDedicatedGateway
//
// RETURNS:
//   - QueryTheDetailsOfTheDedicatedGatewayResponse: The return type of the QueryTheDetailsOfTheDedicatedGateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheDetailsOfTheDedicatedGateway(request *QueryTheDetailsOfTheDedicatedGatewayRequest) (*QueryTheDetailsOfTheDedicatedGatewayResponse, error) {
	result := &QueryTheDetailsOfTheDedicatedGatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheDetailsOfTheDedicatedGatewayUri(VERSION_V1, util.StringValue(request.EtGatewayId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfDedicatedLineGateways
//
// PARAMS:
//   - request: the arguments to QueryTheListOfDedicatedLineGateways
//
// RETURNS:
//   - QueryTheListOfDedicatedLineGatewaysResponse: The return type of the QueryTheListOfDedicatedLineGateways interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfDedicatedLineGateways(request *QueryTheListOfDedicatedLineGatewaysRequest) (*QueryTheListOfDedicatedLineGatewaysResponse, error) {
	result := &QueryTheListOfDedicatedLineGatewaysResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfDedicatedLineGatewaysUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("etGatewayId", util.StringValue(request.EtGatewayId)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfElasticNetworkCards
//
// PARAMS:
//   - request: the arguments to QueryTheListOfElasticNetworkCards
//
// RETURNS:
//   - QueryTheListOfElasticNetworkCardsResponse: The return type of the QueryTheListOfElasticNetworkCards interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfElasticNetworkCards(request *QueryTheListOfElasticNetworkCardsRequest) (*QueryTheListOfElasticNetworkCardsResponse, error) {
	result := &QueryTheListOfElasticNetworkCardsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfElasticNetworkCardsUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("privateIpAddress", strings.Join(util.PtrSliceToStringSlice(request.PrivateIpAddress), ",")).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfEnterpriseSecurityGroups
//
// PARAMS:
//   - request: the arguments to QueryTheListOfEnterpriseSecurityGroups
//
// RETURNS:
//   - QueryTheListOfEnterpriseSecurityGroupsResponse: The return type of the QueryTheListOfEnterpriseSecurityGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfEnterpriseSecurityGroups(request *QueryTheListOfEnterpriseSecurityGroupsRequest) (*QueryTheListOfEnterpriseSecurityGroupsResponse, error) {
	result := &QueryTheListOfEnterpriseSecurityGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfEnterpriseSecurityGroupsUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfHighlyAvailableVirtualIps
//
// PARAMS:
//   - request: the arguments to QueryTheListOfHighlyAvailableVirtualIps
//
// RETURNS:
//   - QueryTheListOfHighlyAvailableVirtualIpsResponse: The return type of the QueryTheListOfHighlyAvailableVirtualIps interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfHighlyAvailableVirtualIps(request *QueryTheListOfHighlyAvailableVirtualIpsRequest) (*QueryTheListOfHighlyAvailableVirtualIpsResponse, error) {
	result := &QueryTheListOfHighlyAvailableVirtualIpsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfHighlyAvailableVirtualIpsUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfIpAddressGroups
//
// PARAMS:
//   - request: the arguments to QueryTheListOfIpAddressGroups
//
// RETURNS:
//   - QueryTheListOfIpAddressGroupsResponse: The return type of the QueryTheListOfIpAddressGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfIpAddressGroups(request *QueryTheListOfIpAddressGroupsRequest) (*QueryTheListOfIpAddressGroupsResponse, error) {
	result := &QueryTheListOfIpAddressGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfIpAddressGroupsUri(VERSION_V1)).
		WithQueryParamFilter("ipVersion", util.StringValue(request.IpVersion)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfPeerConnections
//
// PARAMS:
//   - request: the arguments to QueryTheListOfPeerConnections
//
// RETURNS:
//   - QueryTheListOfPeerConnectionsResponse: The return type of the QueryTheListOfPeerConnections interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfPeerConnections(request *QueryTheListOfPeerConnectionsRequest) (*QueryTheListOfPeerConnectionsResponse, error) {
	result := &QueryTheListOfPeerConnectionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfPeerConnectionsUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfRegularSecurityGroupsV2
//
// PARAMS:
//   - request: the arguments to QueryTheListOfRegularSecurityGroupsV2
//
// RETURNS:
//   - QueryTheListOfRegularSecurityGroupsV2Response: The return type of the QueryTheListOfRegularSecurityGroupsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfRegularSecurityGroupsV2(request *QueryTheListOfRegularSecurityGroupsV2Request) (*QueryTheListOfRegularSecurityGroupsV2Response, error) {
	result := &QueryTheListOfRegularSecurityGroupsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfRegularSecurityGroupsV2Uri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("securityGroupId", util.StringValue(request.SecurityGroupId)).
		WithQueryParamFilter("securityGroupIds", util.StringValue(request.SecurityGroupIds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfSpeedLimitPoliciesForIpv6Gateway
//
// PARAMS:
//   - request: the arguments to QueryTheListOfSpeedLimitPoliciesForIpv6Gateway
//
// RETURNS:
//   - QueryTheListOfSpeedLimitPoliciesForIpv6GatewayResponse: The return type of the QueryTheListOfSpeedLimitPoliciesForIpv6Gateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfSpeedLimitPoliciesForIpv6Gateway(request *QueryTheListOfSpeedLimitPoliciesForIpv6GatewayRequest) (*QueryTheListOfSpeedLimitPoliciesForIpv6GatewayResponse, error) {
	result := &QueryTheListOfSpeedLimitPoliciesForIpv6GatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfSpeedLimitPoliciesForIpv6GatewayUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheSpecifiedElasticNetworkCard
//
// PARAMS:
//   - request: the arguments to QueryTheSpecifiedElasticNetworkCard
//
// RETURNS:
//   - QueryTheSpecifiedElasticNetworkCardResponse: The return type of the QueryTheSpecifiedElasticNetworkCard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheSpecifiedElasticNetworkCard(request *QueryTheSpecifiedElasticNetworkCardRequest) (*QueryTheSpecifiedElasticNetworkCardResponse, error) {
	result := &QueryTheSpecifiedElasticNetworkCardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheSpecifiedElasticNetworkCardUri(VERSION_V1, util.StringValue(request.EniId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheSpecifiedHighlyAvailableVirtualIp
//
// PARAMS:
//   - request: the arguments to QueryTheSpecifiedHighlyAvailableVirtualIp
//
// RETURNS:
//   - QueryTheSpecifiedHighlyAvailableVirtualIpResponse: The return type of the QueryTheSpecifiedHighlyAvailableVirtualIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheSpecifiedHighlyAvailableVirtualIp(request *QueryTheSpecifiedHighlyAvailableVirtualIpRequest) (*QueryTheSpecifiedHighlyAvailableVirtualIpResponse, error) {
	result := &QueryTheSpecifiedHighlyAvailableVirtualIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheSpecifiedHighlyAvailableVirtualIpUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheSpecifiedIpAddressFamily
//
// PARAMS:
//   - request: the arguments to QueryTheSpecifiedIpAddressFamily
//
// RETURNS:
//   - QueryTheSpecifiedIpAddressFamilyResponse: The return type of the QueryTheSpecifiedIpAddressFamily interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheSpecifiedIpAddressFamily(request *QueryTheSpecifiedIpAddressFamilyRequest) (*QueryTheSpecifiedIpAddressFamilyResponse, error) {
	result := &QueryTheSpecifiedIpAddressFamilyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheSpecifiedIpAddressFamilyUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheSpecifiedIpAddressGroup
//
// PARAMS:
//   - request: the arguments to QueryTheSpecifiedIpAddressGroup
//
// RETURNS:
//   - QueryTheSpecifiedIpAddressGroupResponse: The return type of the QueryTheSpecifiedIpAddressGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheSpecifiedIpAddressGroup(request *QueryTheSpecifiedIpAddressGroupRequest) (*QueryTheSpecifiedIpAddressGroupResponse, error) {
	result := &QueryTheSpecifiedIpAddressGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheSpecifiedIpAddressGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheStatusOfTheElasticNetworkCard
//
// PARAMS:
//   - request: the arguments to QueryTheStatusOfTheElasticNetworkCard
//
// RETURNS:
//   - QueryTheStatusOfTheElasticNetworkCardResponse: The return type of the QueryTheStatusOfTheElasticNetworkCard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheStatusOfTheElasticNetworkCard(request *QueryTheStatusOfTheElasticNetworkCardRequest) (*QueryTheStatusOfTheElasticNetworkCardResponse, error) {
	result := &QueryTheStatusOfTheElasticNetworkCardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheStatusOfTheElasticNetworkCardUri(VERSION_V1, util.StringValue(request.EniId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryVpcIntranetIp
//
// PARAMS:
//   - request: the arguments to QueryVpcIntranetIp
//
// RETURNS:
//   - QueryVpcIntranetIpResponse: The return type of the QueryVpcIntranetIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryVpcIntranetIp(request *QueryVpcIntranetIpRequest) (*QueryVpcIntranetIpResponse, error) {
	result := &QueryVpcIntranetIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryVpcIntranetIpUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("privateIpAddresses", strings.Join(util.PtrSliceToStringSlice(request.PrivateIpAddresses), ",")).
		WithQueryParamFilter("privateIpRange", util.StringValue(request.PrivateIpRange)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryVpcList
//
// PARAMS:
//   - request: the arguments to QueryVpcList
//
// RETURNS:
//   - QueryVpcListResponse: The return type of the QueryVpcList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryVpcList(request *QueryVpcListRequest) (*QueryVpcListResponse, error) {
	result := &QueryVpcListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryVpcListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("isDefault", util.BoolValue(request.IsDefault)).
		WithQueryParamFilter("vpcIds", util.StringValue(request.VpcIds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryVpnList
//
// PARAMS:
//   - request: the arguments to QueryVpnList
//
// RETURNS:
//   - QueryVpnListResponse: The return type of the QueryVpnList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryVpnList(request *QueryVpnListRequest) (*QueryVpnListResponse, error) {
	result := &QueryVpnListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryVpnListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("eip", util.StringValue(request.Eip)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusion
//
// PARAMS:
//   - request: the arguments to QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusion
//
// RETURNS:
//   - QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionResponse: The return type of the QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusion(request *QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionRequest) (*QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionResponse, error) {
	result := &QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RejectPeerToPeerConnectionRequest
//
// PARAMS:
//   - request: the arguments to RejectPeerToPeerConnectionRequest
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RejectPeerToPeerConnectionRequest(request *RejectPeerToPeerConnectionRequestRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRejectPeerToPeerConnectionRequestUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("reject", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleaseDedicatedGateway
//
// PARAMS:
//   - request: the arguments to ReleaseDedicatedGateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseDedicatedGateway(request *ReleaseDedicatedGatewayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseDedicatedGatewayUri(VERSION_V1, util.StringValue(request.EtGatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleasePeerToPeerConnection
//
// PARAMS:
//   - request: the arguments to ReleasePeerToPeerConnection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleasePeerToPeerConnection(request *ReleasePeerToPeerConnectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleasePeerToPeerConnectionUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleaseVpn
//
// PARAMS:
//   - request: the arguments to ReleaseVpn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseVpn(request *ReleaseVpnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseVpnUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RemoveElasticNetworkCard
//
// PARAMS:
//   - request: the arguments to RemoveElasticNetworkCard
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveElasticNetworkCard(request *RemoveElasticNetworkCardRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveElasticNetworkCardUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RemoveIpAddressGroupFromIpAddressFamily
//
// PARAMS:
//   - request: the arguments to RemoveIpAddressGroupFromIpAddressFamily
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveIpAddressGroupFromIpAddressFamily(request *RemoveIpAddressGroupFromIpAddressFamilyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveIpAddressGroupFromIpAddressFamilyUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RenewVpn
//
// PARAMS:
//   - request: the arguments to RenewVpn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenewVpn(request *RenewVpnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenewVpnUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RevokeRegularSecurityGroupRulesV2
//
// PARAMS:
//   - request: the arguments to RevokeRegularSecurityGroupRulesV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RevokeRegularSecurityGroupRulesV2(request *RevokeRegularSecurityGroupRulesV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRevokeRegularSecurityGroupRulesV2Uri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		WithQueryParam("revokeRule", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		WithBody(request).
		Do()
}

// SearchForVpnDetails
//
// PARAMS:
//   - request: the arguments to SearchForVpnDetails
//
// RETURNS:
//   - SearchForVpnDetailsResponse: The return type of the SearchForVpnDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SearchForVpnDetails(request *SearchForVpnDetailsRequest) (*SearchForVpnDetailsResponse, error) {
	result := &SearchForVpnDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSearchForVpnDetailsUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SearchVpnTunnel
//
// PARAMS:
//   - request: the arguments to SearchVpnTunnel
//
// RETURNS:
//   - SearchVpnTunnelResponse: The return type of the SearchVpnTunnel interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SearchVpnTunnel(request *SearchVpnTunnelRequest) (*SearchVpnTunnelResponse, error) {
	result := &SearchVpnTunnelResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSearchVpnTunnelUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindEip
//
// PARAMS:
//   - request: the arguments to UnbindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindEip(request *UnbindEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindEipUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UnbindPhysicalDedicatedLine
//
// PARAMS:
//   - request: the arguments to UnbindPhysicalDedicatedLine
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindPhysicalDedicatedLine(request *UnbindPhysicalDedicatedLineRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindPhysicalDedicatedLineUri(VERSION_V1, util.StringValue(request.EtGatewayId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UpdateAclRules
//
// PARAMS:
//   - request: the arguments to UpdateAclRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAclRules(request *UpdateAclRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAclRulesUri(VERSION_V1, util.StringValue(request.AclRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateDedicatedGateway
//
// PARAMS:
//   - request: the arguments to UpdateDedicatedGateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDedicatedGateway(request *UpdateDedicatedGatewayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDedicatedGatewayUri(VERSION_V1, util.StringValue(request.EtGatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateElasticNetworkCard
//
// PARAMS:
//   - request: the arguments to UpdateElasticNetworkCard
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateElasticNetworkCard(request *UpdateElasticNetworkCardRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateElasticNetworkCardUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEnterpriseSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to UpdateEnterpriseSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEnterpriseSecurityGroupRules(request *UpdateEnterpriseSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEnterpriseSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.EnterpriseSecurityGroupRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateHighlyAvailableVirtualIp
//
// PARAMS:
//   - request: the arguments to UpdateHighlyAvailableVirtualIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateHighlyAvailableVirtualIp(request *UpdateHighlyAvailableVirtualIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateHighlyAvailableVirtualIpUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateIpAddressFamily
//
// PARAMS:
//   - request: the arguments to UpdateIpAddressFamily
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIpAddressFamily(request *UpdateIpAddressFamilyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIpAddressFamilyUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateIpAddressGroup
//
// PARAMS:
//   - request: the arguments to UpdateIpAddressGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIpAddressGroup(request *UpdateIpAddressGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIpAddressGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateIpv6GatewayReleaseProtectionSwitch
//
// PARAMS:
//   - request: the arguments to UpdateIpv6GatewayReleaseProtectionSwitch
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIpv6GatewayReleaseProtectionSwitch(request *UpdateIpv6GatewayReleaseProtectionSwitchRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIpv6GatewayReleaseProtectionSwitchUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateIpv6GatewaySpeedLimitPolicy
//
// PARAMS:
//   - request: the arguments to UpdateIpv6GatewaySpeedLimitPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIpv6GatewaySpeedLimitPolicy(request *UpdateIpv6GatewaySpeedLimitPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIpv6GatewaySpeedLimitPolicyUri(VERSION_V1, util.StringValue(request.GatewayId), util.StringValue(request.RateLimitRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateNetworkDetection
//
// PARAMS:
//   - request: the arguments to UpdateNetworkDetection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateNetworkDetection(request *UpdateNetworkDetectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateNetworkDetectionUri(VERSION_V1, util.StringValue(request.ProbeId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdatePeerToPeerConnectionReleaseProtectionSwitch
//
// PARAMS:
//   - request: the arguments to UpdatePeerToPeerConnectionReleaseProtectionSwitch
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePeerToPeerConnectionReleaseProtectionSwitch(request *UpdatePeerToPeerConnectionReleaseProtectionSwitchRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePeerToPeerConnectionReleaseProtectionSwitchUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithBody(request).
		Do()
}

// UpdateRegularSecurityGroupRulesV2
//
// PARAMS:
//   - request: the arguments to UpdateRegularSecurityGroupRulesV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRegularSecurityGroupRulesV2(request *UpdateRegularSecurityGroupRulesV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRegularSecurityGroupRulesV2Uri(VERSION_V2)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		WithBody(request).
		Do()
}

// UpdateRoutingRules
//
// PARAMS:
//   - request: the arguments to UpdateRoutingRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRoutingRules(request *UpdateRoutingRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRoutingRulesUri(VERSION_V1, util.StringValue(request.RouteRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateSslVpnServer
//
// PARAMS:
//   - request: the arguments to UpdateSslVpnServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSslVpnServer(request *UpdateSslVpnServerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSslVpnServerUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.SslVpnServerId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateSslVpnUsers
//
// PARAMS:
//   - request: the arguments to UpdateSslVpnUsers
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSslVpnUsers(request *UpdateSslVpnUsersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSslVpnUsersUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.UserId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateSubnet
//
// PARAMS:
//   - request: the arguments to UpdateSubnet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSubnet(request *UpdateSubnetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnections
//
// PARAMS:
//   - request: the arguments to UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnections
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnections(request *UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithBody(request).
		Do()
}

// UpdateUserGateway
//
// PARAMS:
//   - request: the arguments to UpdateUserGateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateUserGateway(request *UpdateUserGatewayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateUserGatewayUri(VERSION_V1, util.StringValue(request.CgwId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateVpc
//
// PARAMS:
//   - request: the arguments to UpdateVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpc(request *UpdateVpcRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateVpn
//
// PARAMS:
//   - request: the arguments to UpdateVpn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpn(request *UpdateVpnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpnUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateVpnReleaseProtection
//
// PARAMS:
//   - request: the arguments to UpdateVpnReleaseProtection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpnReleaseProtection(request *UpdateVpnReleaseProtectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpnReleaseProtectionUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateVpnTunnel
//
// PARAMS:
//   - request: the arguments to UpdateVpnTunnel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpnTunnel(request *UpdateVpnTunnelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpnTunnelUri(VERSION_V1, util.StringValue(request.VpnConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UserGatewayDetails
//
// PARAMS:
//   - request: the arguments to UserGatewayDetails
//
// RETURNS:
//   - UserGatewayDetailsResponse: The return type of the UserGatewayDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserGatewayDetails(request *UserGatewayDetailsRequest) (*UserGatewayDetailsResponse, error) {
	result := &UserGatewayDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getUserGatewayDetailsUri(VERSION_V1, util.StringValue(request.CgwId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserGatewayList
//
// PARAMS:
//   - request: the arguments to UserGatewayList
//
// RETURNS:
//   - UserGatewayListResponse: The return type of the UserGatewayList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserGatewayList(request *UserGatewayListRequest) (*UserGatewayListResponse, error) {
	result := &UserGatewayListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getUserGatewayListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ViewGatewayLimitRules
//
// PARAMS:
//   - request: the arguments to ViewGatewayLimitRules
//
// RETURNS:
//   - ViewGatewayLimitRulesResponse: The return type of the ViewGatewayLimitRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ViewGatewayLimitRules(request *ViewGatewayLimitRulesRequest) (*ViewGatewayLimitRulesResponse, error) {
	result := &ViewGatewayLimitRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getViewGatewayLimitRulesUri(VERSION_V1)).
		WithQueryParamFilter("serviceType", util.StringValue(request.ServiceType)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("glrId", util.StringValue(request.GlrId)).
		WithQueryParamFilter("resourceId", util.StringValue(request.ResourceId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.StringValue(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ViewPeerToPeerConnectionDetails
//
// PARAMS:
//   - request: the arguments to ViewPeerToPeerConnectionDetails
//
// RETURNS:
//   - ViewPeerToPeerConnectionDetailsResponse: The return type of the ViewPeerToPeerConnectionDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ViewPeerToPeerConnectionDetails(request *ViewPeerToPeerConnectionDetailsRequest) (*ViewPeerToPeerConnectionDetailsResponse, error) {
	result := &ViewPeerToPeerConnectionDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getViewPeerToPeerConnectionDetailsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParamFilter("role", util.StringValue(request.Role)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ViewSecurityGroupDetailsV2
//
// PARAMS:
//   - request: the arguments to ViewSecurityGroupDetailsV2
//
// RETURNS:
//   - ViewSecurityGroupDetailsV2Response: The return type of the ViewSecurityGroupDetailsV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ViewSecurityGroupDetailsV2(request *ViewSecurityGroupDetailsV2Request) (*ViewSecurityGroupDetailsV2Response, error) {
	result := &ViewSecurityGroupDetailsV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getViewSecurityGroupDetailsV2Uri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
