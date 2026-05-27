package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"strings"
)

const (
	VERSION_V1 = "v1"
)

// AcceptPeerConn
//
// PARAMS:
//   - request: the arguments to AcceptPeerConn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AcceptPeerConn(request *AcceptPeerConnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAcceptPeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
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

// AddEniIp
//
// PARAMS:
//   - request: the arguments to AddEniIp
//
// RETURNS:
//   - AddEniIpResponse: The return type of the AddEniIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddEniIp(request *AddEniIpRequest) (*AddEniIpResponse, error) {
	result := &AddEniIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddEniIpUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddIpAddressToIpGroup
//
// PARAMS:
//   - request: the arguments to AddIpAddressToIpGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddIpAddressToIpGroup(request *AddIpAddressToIpGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddIpAddressToIpGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddIpGroupToIpSet
//
// PARAMS:
//   - request: the arguments to AddIpGroupToIpSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddIpGroupToIpSet(request *AddIpGroupToIpSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddIpGroupToIpSetUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AttachEniInstance
//
// PARAMS:
//   - request: the arguments to AttachEniInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AttachEniInstance(request *AttachEniInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAttachEniInstanceUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("attach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AuthorizeEnterpriseSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to AuthorizeEnterpriseSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuthorizeEnterpriseSecurityGroupRules(request *AuthorizeEnterpriseSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAuthorizeEnterpriseSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.EnterpriseSecurityGroupId))).
		WithQueryParam("authorizeRule", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AuthorizeSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to AuthorizeSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuthorizeSecurityGroupRules(request *AuthorizeSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAuthorizeSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.SecurityGroupId))).
		WithQueryParam("authorizeRule", "").
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BatchAddDnatRules
//
// PARAMS:
//   - request: the arguments to BatchAddDnatRules
//
// RETURNS:
//   - BatchAddDnatRulesResponse: The return type of the BatchAddDnatRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchAddDnatRules(request *BatchAddDnatRulesRequest) (*BatchAddDnatRulesResponse, error) {
	result := &BatchAddDnatRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchAddDnatRulesUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchAddEniIp
//
// PARAMS:
//   - request: the arguments to BatchAddEniIp
//
// RETURNS:
//   - BatchAddEniIpResponse: The return type of the BatchAddEniIp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchAddEniIp(request *BatchAddEniIpRequest) (*BatchAddEniIpResponse, error) {
	result := &BatchAddEniIpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchAddEniIpUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchAddSnatRules
//
// PARAMS:
//   - request: the arguments to BatchAddSnatRules
//
// RETURNS:
//   - BatchAddSnatRulesResponse: The return type of the BatchAddSnatRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchAddSnatRules(request *BatchAddSnatRulesRequest) (*BatchAddSnatRulesResponse, error) {
	result := &BatchAddSnatRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchAddSnatRulesUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// BatchDeleteEniIp
//
// PARAMS:
//   - request: the arguments to BatchDeleteEniIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteEniIp(request *BatchDeleteEniIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchDeleteEniIpUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
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

// BindEniEip
//
// PARAMS:
//   - request: the arguments to BindEniEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindEniEip(request *BindEniEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindEniEipUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindHaVipEip
//
// PARAMS:
//   - request: the arguments to BindHaVipEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindHaVipEip(request *BindHaVipEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindHaVipEipUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("bindPublicIp", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindHaVipInstance
//
// PARAMS:
//   - request: the arguments to BindHaVipInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindHaVipInstance(request *BindHaVipInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindHaVipInstanceUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("attach", "").
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

// ClosePeerConnSyncDns
//
// PARAMS:
//   - request: the arguments to ClosePeerConnSyncDns
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ClosePeerConnSyncDns(request *ClosePeerConnSyncDnsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClosePeerConnSyncDnsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
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

// CreateDnatRule
//
// PARAMS:
//   - request: the arguments to CreateDnatRule
//
// RETURNS:
//   - CreateDnatRuleResponse: The return type of the CreateDnatRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDnatRule(request *CreateDnatRuleRequest) (*CreateDnatRuleResponse, error) {
	result := &CreateDnatRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDnatRuleUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEgressOnlyRule
//
// PARAMS:
//   - request: the arguments to CreateEgressOnlyRule
//
// RETURNS:
//   - CreateEgressOnlyRuleResponse: The return type of the CreateEgressOnlyRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEgressOnlyRule(request *CreateEgressOnlyRuleRequest) (*CreateEgressOnlyRuleResponse, error) {
	result := &CreateEgressOnlyRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEgressOnlyRuleUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEni
//
// PARAMS:
//   - request: the arguments to CreateEni
//
// RETURNS:
//   - CreateEniResponse: The return type of the CreateEni interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEni(request *CreateEniRequest) (*CreateEniResponse, error) {
	result := &CreateEniResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEniUri(VERSION_V1)).
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

// CreateHaVip
//
// PARAMS:
//   - request: the arguments to CreateHaVip
//
// RETURNS:
//   - CreateHaVipResponse: The return type of the CreateHaVip interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (*CreateHaVipResponse, error) {
	result := &CreateHaVipResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateHaVipUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIpGroup
//
// PARAMS:
//   - request: the arguments to CreateIpGroup
//
// RETURNS:
//   - CreateIpGroupResponse: The return type of the CreateIpGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpGroup(request *CreateIpGroupRequest) (*CreateIpGroupResponse, error) {
	result := &CreateIpGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpGroupUri(VERSION_V1)).
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

// CreateIpSet
//
// PARAMS:
//   - request: the arguments to CreateIpSet
//
// RETURNS:
//   - CreateIpSetResponse: The return type of the CreateIpSet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpSet(request *CreateIpSetRequest) (*CreateIpSetResponse, error) {
	result := &CreateIpSetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpSetUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIpv6Gateway
//
// PARAMS:
//   - request: the arguments to CreateIpv6Gateway
//
// RETURNS:
//   - CreateIpv6GatewayResponse: The return type of the CreateIpv6Gateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIpv6Gateway(request *CreateIpv6GatewayRequest) (*CreateIpv6GatewayResponse, error) {
	result := &CreateIpv6GatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIpv6GatewayUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateNat
//
// PARAMS:
//   - request: the arguments to CreateNat
//
// RETURNS:
//   - CreateNatResponse: The return type of the CreateNat interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateNat(request *CreateNatRequest) (*CreateNatResponse, error) {
	result := &CreateNatResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateNatUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePeerConn
//
// PARAMS:
//   - request: the arguments to CreatePeerConn
//
// RETURNS:
//   - CreatePeerConnResponse: The return type of the CreatePeerConn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePeerConn(request *CreatePeerConnRequest) (*CreatePeerConnResponse, error) {
	result := &CreatePeerConnResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePeerConnUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateProbe
//
// PARAMS:
//   - request: the arguments to CreateProbe
//
// RETURNS:
//   - CreateProbeResponse: The return type of the CreateProbe interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateProbe(request *CreateProbeRequest) (*CreateProbeResponse, error) {
	result := &CreateProbeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateProbeUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRateLimitRule
//
// PARAMS:
//   - request: the arguments to CreateRateLimitRule
//
// RETURNS:
//   - CreateRateLimitRuleResponse: The return type of the CreateRateLimitRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateRateLimitRule(request *CreateRateLimitRuleRequest) (*CreateRateLimitRuleResponse, error) {
	result := &CreateRateLimitRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRateLimitRuleUri(VERSION_V1, util.StringValue(request.GatewayId))).
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

// CreateSecurityGroup
//
// PARAMS:
//   - request: the arguments to CreateSecurityGroup
//
// RETURNS:
//   - CreateSecurityGroupResponse: The return type of the CreateSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	result := &CreateSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSecurityGroupUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSnatRule
//
// PARAMS:
//   - request: the arguments to CreateSnatRule
//
// RETURNS:
//   - CreateSnatRuleResponse: The return type of the CreateSnatRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSnatRule(request *CreateSnatRuleRequest) (*CreateSnatRuleResponse, error) {
	result := &CreateSnatRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSnatRuleUri(VERSION_V1, util.StringValue(request.NatId))).
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

// DeleteDnatRule
//
// PARAMS:
//   - request: the arguments to DeleteDnatRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDnatRule(request *DeleteDnatRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDnatRuleUri(VERSION_V1, util.StringValue(request.NatId), util.StringValue(request.RuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteEniIp
//
// PARAMS:
//   - request: the arguments to DeleteEniIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteEniIp(request *DeleteEniIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteEniIpUri(VERSION_V1, util.StringValue(request.EniId), util.StringValue(request.PrivateIpAddress))).
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

// DeleteHaVip
//
// PARAMS:
//   - request: the arguments to DeleteHaVip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteHaVipUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpGroup
//
// PARAMS:
//   - request: the arguments to DeleteIpGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpGroup(request *DeleteIpGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
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

// DeleteIpSet
//
// PARAMS:
//   - request: the arguments to DeleteIpSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpSet(request *DeleteIpSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpSetUri(VERSION_V1, util.StringValue(request.IpGroupId))).
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

// DeleteIpv6GatewayEgressOnlyRule
//
// PARAMS:
//   - request: the arguments to DeleteIpv6GatewayEgressOnlyRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpv6GatewayEgressOnlyRule(request *DeleteIpv6GatewayEgressOnlyRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpv6GatewayEgressOnlyRuleUri(VERSION_V1, util.StringValue(request.GatewayId), util.StringValue(request.EgressOnlyRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteIpv6GatewayRateLimitRule
//
// PARAMS:
//   - request: the arguments to DeleteIpv6GatewayRateLimitRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpv6GatewayRateLimitRule(request *DeleteIpv6GatewayRateLimitRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpv6GatewayRateLimitRuleUri(VERSION_V1, util.StringValue(request.GatewayId), util.StringValue(request.RateLimitRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteProbe
//
// PARAMS:
//   - request: the arguments to DeleteProbe
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteProbe(request *DeleteProbeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteProbeUri(VERSION_V1, util.StringValue(request.ProbeId))).
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

// DeleteSecurityGroup
//
// PARAMS:
//   - request: the arguments to DeleteSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSecurityGroupUri(VERSION_V1, util.StringValue(request.SecurityGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to DeleteSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityGroupRules(request *DeleteSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.SecurityGroupRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		Do()
}

// DeleteSnatRule
//
// PARAMS:
//   - request: the arguments to DeleteSnatRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSnatRule(request *DeleteSnatRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSnatRuleUri(VERSION_V1, util.StringValue(request.NatId), util.StringValue(request.RuleId))).
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

// DetachEniInstance
//
// PARAMS:
//   - request: the arguments to DetachEniInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DetachEniInstance(request *DetachEniInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDetachEniInstanceUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("detach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// GetEniDetail
//
// PARAMS:
//   - request: the arguments to GetEniDetail
//
// RETURNS:
//   - GetEniDetailResponse: The return type of the GetEniDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetEniDetail(request *GetEniDetailRequest) (*GetEniDetailResponse, error) {
	result := &GetEniDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetEniDetailUri(VERSION_V1, util.StringValue(request.EniId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEniStatus
//
// PARAMS:
//   - request: the arguments to GetEniStatus
//
// RETURNS:
//   - GetEniStatusResponse: The return type of the GetEniStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetEniStatus(request *GetEniStatusRequest) (*GetEniStatusResponse, error) {
	result := &GetEniStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetEniStatusUri(VERSION_V1, util.StringValue(request.EniId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHaVipDetail
//
// PARAMS:
//   - request: the arguments to GetHaVipDetail
//
// RETURNS:
//   - GetHaVipDetailResponse: The return type of the GetHaVipDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetHaVipDetail(request *GetHaVipDetailRequest) (*GetHaVipDetailResponse, error) {
	result := &GetHaVipDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetHaVipDetailUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNat
//
// PARAMS:
//   - request: the arguments to GetNat
//
// RETURNS:
//   - GetNatResponse: The return type of the GetNat interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetNat(request *GetNatRequest) (*GetNatResponse, error) {
	result := &GetNatResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetNatUri(VERSION_V1, util.StringValue(request.NatId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPeerConn
//
// PARAMS:
//   - request: the arguments to GetPeerConn
//
// RETURNS:
//   - GetPeerConnResponse: The return type of the GetPeerConn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPeerConn(request *GetPeerConnRequest) (*GetPeerConnResponse, error) {
	result := &GetPeerConnResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParamFilter("role", util.StringValue(request.Role)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetProbeDetail
//
// PARAMS:
//   - request: the arguments to GetProbeDetail
//
// RETURNS:
//   - GetProbeDetailResponse: The return type of the GetProbeDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetProbeDetail(request *GetProbeDetailRequest) (*GetProbeDetailResponse, error) {
	result := &GetProbeDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetProbeDetailUri(VERSION_V1, util.StringValue(request.ProbeId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSecurityGroupDetails
//
// PARAMS:
//   - request: the arguments to GetSecurityGroupDetails
//
// RETURNS:
//   - GetSecurityGroupDetailsResponse: The return type of the GetSecurityGroupDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSecurityGroupDetails(request *GetSecurityGroupDetailsRequest) (*GetSecurityGroupDetailsResponse, error) {
	result := &GetSecurityGroupDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSecurityGroupDetailsUri(VERSION_V1, util.StringValue(request.SecurityGroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// ListDnatRule
//
// PARAMS:
//   - request: the arguments to ListDnatRule
//
// RETURNS:
//   - ListDnatRuleResponse: The return type of the ListDnatRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListDnatRule(request *ListDnatRuleRequest) (*ListDnatRuleResponse, error) {
	result := &ListDnatRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListDnatRuleUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListEgressOnlyRule
//
// PARAMS:
//   - request: the arguments to ListEgressOnlyRule
//
// RETURNS:
//   - ListEgressOnlyRuleResponse: The return type of the ListEgressOnlyRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListEgressOnlyRule(request *ListEgressOnlyRuleRequest) (*ListEgressOnlyRuleResponse, error) {
	result := &ListEgressOnlyRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListEgressOnlyRuleUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListEni
//
// PARAMS:
//   - request: the arguments to ListEni
//
// RETURNS:
//   - ListEniResponse: The return type of the ListEni interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListEni(request *ListEniRequest) (*ListEniResponse, error) {
	result := &ListEniResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListEniUri(VERSION_V1)).
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

// ListHaVip
//
// PARAMS:
//   - request: the arguments to ListHaVip
//
// RETURNS:
//   - ListHaVipResponse: The return type of the ListHaVip interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListHaVip(request *ListHaVipRequest) (*ListHaVipResponse, error) {
	result := &ListHaVipResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListHaVipUri(VERSION_V1)).
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

// ListNat
//
// PARAMS:
//   - request: the arguments to ListNat
//
// RETURNS:
//   - ListNatResponse: The return type of the ListNat interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListNat(request *ListNatRequest) (*ListNatResponse, error) {
	result := &ListNatResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListNatUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("natId", util.StringValue(request.NatId)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPeerConn
//
// PARAMS:
//   - request: the arguments to ListPeerConn
//
// RETURNS:
//   - ListPeerConnResponse: The return type of the ListPeerConn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPeerConn(request *ListPeerConnRequest) (*ListPeerConnResponse, error) {
	result := &ListPeerConnResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPeerConnUri(VERSION_V1)).
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

// ListProbes
//
// PARAMS:
//   - request: the arguments to ListProbes
//
// RETURNS:
//   - ListProbesResponse: The return type of the ListProbes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListProbes(request *ListProbesRequest) (*ListProbesResponse, error) {
	result := &ListProbesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListProbesUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRateLimitRule
//
// PARAMS:
//   - request: the arguments to ListRateLimitRule
//
// RETURNS:
//   - ListRateLimitRuleResponse: The return type of the ListRateLimitRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRateLimitRule(request *ListRateLimitRuleRequest) (*ListRateLimitRuleResponse, error) {
	result := &ListRateLimitRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRateLimitRuleUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSnatRule
//
// PARAMS:
//   - request: the arguments to ListSnatRule
//
// RETURNS:
//   - ListSnatRuleResponse: The return type of the ListSnatRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSnatRule(request *ListSnatRuleRequest) (*ListSnatRuleResponse, error) {
	result := &ListSnatRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSnatRuleUri(VERSION_V1, util.StringValue(request.NatId))).
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

// ModifyNat
//
// PARAMS:
//   - request: the arguments to ModifyNat
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyNat(request *ModifyNatRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyNatUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// NatBindEip
//
// PARAMS:
//   - request: the arguments to NatBindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) NatBindEip(request *NatBindEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getNatBindEipUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// NatUnBindEip
//
// PARAMS:
//   - request: the arguments to NatUnBindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) NatUnBindEip(request *NatUnBindEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getNatUnBindEipUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// OpenPeerConnSyncDns
//
// PARAMS:
//   - request: the arguments to OpenPeerConnSyncDns
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) OpenPeerConnSyncDns(request *OpenPeerConnSyncDnsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getOpenPeerConnSyncDnsUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("open", "").
		WithQueryParamFilter("role", util.StringValue(request.Role)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
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

// PurchaseReservedNat
//
// PARAMS:
//   - request: the arguments to PurchaseReservedNat
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedNat(request *PurchaseReservedNatRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPurchaseReservedNatUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
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

// QueryEnterpriseSecurityGroupList
//
// PARAMS:
//   - request: the arguments to QueryEnterpriseSecurityGroupList
//
// RETURNS:
//   - QueryEnterpriseSecurityGroupListResponse: The return type of the QueryEnterpriseSecurityGroupList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryEnterpriseSecurityGroupList(request *QueryEnterpriseSecurityGroupListRequest) (*QueryEnterpriseSecurityGroupListResponse, error) {
	result := &QueryEnterpriseSecurityGroupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryEnterpriseSecurityGroupListUri(VERSION_V1)).
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

// QueryIpGroupDetail
//
// PARAMS:
//   - request: the arguments to QueryIpGroupDetail
//
// RETURNS:
//   - QueryIpGroupDetailResponse: The return type of the QueryIpGroupDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpGroupDetail(request *QueryIpGroupDetailRequest) (*QueryIpGroupDetailResponse, error) {
	result := &QueryIpGroupDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpGroupDetailUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryIpGroupList
//
// PARAMS:
//   - request: the arguments to QueryIpGroupList
//
// RETURNS:
//   - QueryIpGroupListResponse: The return type of the QueryIpGroupList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpGroupList(request *QueryIpGroupListRequest) (*QueryIpGroupListResponse, error) {
	result := &QueryIpGroupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpGroupListUri(VERSION_V1)).
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

// QueryIpSetDetail
//
// PARAMS:
//   - request: the arguments to QueryIpSetDetail
//
// RETURNS:
//   - QueryIpSetDetailResponse: The return type of the QueryIpSetDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpSetDetail(request *QueryIpSetDetailRequest) (*QueryIpSetDetailResponse, error) {
	result := &QueryIpSetDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpSetDetailUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryIpSetList
//
// PARAMS:
//   - request: the arguments to QueryIpSetList
//
// RETURNS:
//   - QueryIpSetListResponse: The return type of the QueryIpSetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpSetList(request *QueryIpSetListRequest) (*QueryIpSetListResponse, error) {
	result := &QueryIpSetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpSetListUri(VERSION_V1)).
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

// QuerySecurityGroupsList
//
// PARAMS:
//   - request: the arguments to QuerySecurityGroupsList
//
// RETURNS:
//   - QuerySecurityGroupsListResponse: The return type of the QuerySecurityGroupsList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySecurityGroupsList(request *QuerySecurityGroupsListRequest) (*QuerySecurityGroupsListResponse, error) {
	result := &QuerySecurityGroupsListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySecurityGroupsListUri(VERSION_V1)).
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

// RefundPeerConn
//
// PARAMS:
//   - request: the arguments to RefundPeerConn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RefundPeerConn(request *RefundPeerConnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRefundPeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("refund", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RejectPeerConn
//
// PARAMS:
//   - request: the arguments to RejectPeerConn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RejectPeerConn(request *RejectPeerConnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRejectPeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
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

// ReleaseNat
//
// PARAMS:
//   - request: the arguments to ReleaseNat
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseNat(request *ReleaseNatRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseNatUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleasePeerConn
//
// PARAMS:
//   - request: the arguments to ReleasePeerConn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleasePeerConn(request *ReleasePeerConnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleasePeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
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

// RemoveEni
//
// PARAMS:
//   - request: the arguments to RemoveEni
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveEni(request *RemoveEniRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveEniUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RemoveIpAddressFromIpGroup
//
// PARAMS:
//   - request: the arguments to RemoveIpAddressFromIpGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveIpAddressFromIpGroup(request *RemoveIpAddressFromIpGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveIpAddressFromIpGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RemoveIpGroupFromIpSet
//
// PARAMS:
//   - request: the arguments to RemoveIpGroupFromIpSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveIpGroupFromIpSet(request *RemoveIpGroupFromIpSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveIpGroupFromIpSetUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RenewPeerConn
//
// PARAMS:
//   - request: the arguments to RenewPeerConn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenewPeerConn(request *RenewPeerConnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenewPeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("purchaseReserved", "").
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

// ResizeIpv6Gateway
//
// PARAMS:
//   - request: the arguments to ResizeIpv6Gateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeIpv6Gateway(request *ResizeIpv6GatewayRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeIpv6GatewayUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ResizeNat
//
// PARAMS:
//   - request: the arguments to ResizeNat
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeNat(request *ResizeNatRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeNatUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RevokeSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to RevokeSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RevokeSecurityGroupRules(request *RevokeSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRevokeSecurityGroupRulesUri(VERSION_V1, util.StringValue(request.SecurityGroupId))).
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

// UnbindEniEip
//
// PARAMS:
//   - request: the arguments to UnbindEniEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindEniEip(request *UnbindEniEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindEniEipUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("unBind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UnbindHaVipEip
//
// PARAMS:
//   - request: the arguments to UnbindHaVipEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindHaVipEip(request *UnbindHaVipEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindHaVipEipUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("unbindPublicIp", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UnbindHaVipInstance
//
// PARAMS:
//   - request: the arguments to UnbindHaVipInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindHaVipInstance(request *UnbindHaVipInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindHaVipInstanceUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("detach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
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

// UpdateDeleteProtect
//
// PARAMS:
//   - request: the arguments to UpdateDeleteProtect
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDeleteProtect(request *UpdateDeleteProtectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDeleteProtectUri(VERSION_V1, util.StringValue(request.GatewayId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateDnatRule
//
// PARAMS:
//   - request: the arguments to UpdateDnatRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDnatRule(request *UpdateDnatRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDnatRuleUri(VERSION_V1, util.StringValue(request.NatId), util.StringValue(request.RuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEni
//
// PARAMS:
//   - request: the arguments to UpdateEni
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEni(request *UpdateEniRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEniUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEniEnterpriseSecurityGroup
//
// PARAMS:
//   - request: the arguments to UpdateEniEnterpriseSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEniEnterpriseSecurityGroup(request *UpdateEniEnterpriseSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEniEnterpriseSecurityGroupUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("bindEsg", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEniSecurityGroup
//
// PARAMS:
//   - request: the arguments to UpdateEniSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEniSecurityGroup(request *UpdateEniSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEniSecurityGroupUri(VERSION_V1, util.StringValue(request.EniId))).
		WithQueryParam("bindSg", "").
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

// UpdateHaVip
//
// PARAMS:
//   - request: the arguments to UpdateHaVip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateHaVip(request *UpdateHaVipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateHaVipUri(VERSION_V1, util.StringValue(request.HaVipId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateIpGroup
//
// PARAMS:
//   - request: the arguments to UpdateIpGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIpGroup(request *UpdateIpGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIpGroupUri(VERSION_V1, util.StringValue(request.IpSetId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateIpSet
//
// PARAMS:
//   - request: the arguments to UpdateIpSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateIpSet(request *UpdateIpSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateIpSetUri(VERSION_V1, util.StringValue(request.IpGroupId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateNatReleaseProtectionSwitch
//
// PARAMS:
//   - request: the arguments to UpdateNatReleaseProtectionSwitch
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateNatReleaseProtectionSwitch(request *UpdateNatReleaseProtectionSwitchRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateNatReleaseProtectionSwitchUri(VERSION_V1, util.StringValue(request.NatId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdatePeerConn
//
// PARAMS:
//   - request: the arguments to UpdatePeerConn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePeerConn(request *UpdatePeerConnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePeerConnUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdatePeerConnBandwidth
//
// PARAMS:
//   - request: the arguments to UpdatePeerConnBandwidth
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePeerConnBandwidth(request *UpdatePeerConnBandwidthRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePeerConnBandwidthUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdatePeerConnDeleteProtect
//
// PARAMS:
//   - request: the arguments to UpdatePeerConnDeleteProtect
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePeerConnDeleteProtect(request *UpdatePeerConnDeleteProtectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePeerConnDeleteProtectUri(VERSION_V1, util.StringValue(request.PeerConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateProbe
//
// PARAMS:
//   - request: the arguments to UpdateProbe
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateProbe(request *UpdateProbeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateProbeUri(VERSION_V1, util.StringValue(request.ProbeId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateRateLimitRule
//
// PARAMS:
//   - request: the arguments to UpdateRateLimitRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRateLimitRule(request *UpdateRateLimitRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRateLimitRuleUri(VERSION_V1, util.StringValue(request.GatewayId), util.StringValue(request.RateLimitRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
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

// UpdateSecurityGroupRules
//
// PARAMS:
//   - request: the arguments to UpdateSecurityGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityGroupRules(request *UpdateSecurityGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSecurityGroupRulesUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("sgVersion", util.Int64Value(request.SgVersion)).
		WithBody(request).
		Do()
}

// UpdateSnatRule
//
// PARAMS:
//   - request: the arguments to UpdateSnatRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSnatRule(request *UpdateSnatRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSnatRuleUri(VERSION_V1, util.StringValue(request.NatId), util.StringValue(request.RuleId))).
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
