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
