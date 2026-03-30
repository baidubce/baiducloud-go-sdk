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
	return result, err
}

// BindEip
//
// PARAMS:
//   - request: the arguments to BindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindEip(request *BindEipRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindEipUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// CloseVpcRelay
//
// PARAMS:
//   - request: the arguments to CloseVpcRelay
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CloseVpcRelay(request *CloseVpcRelayRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCloseVpcRelayUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
}

// DeleteGatewayLimitRule
//
// PARAMS:
//   - request: the arguments to DeleteGatewayLimitRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteGatewayLimitRule(request *DeleteGatewayLimitRuleRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteGatewayLimitRuleUri(VERSION_V1, util.StringValue(request.GlrId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeleteIpReserve
//
// PARAMS:
//   - request: the arguments to DeleteIpReserve
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpReserve(request *DeleteIpReserveRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteIpReserveUri(VERSION_V1, util.StringValue(request.IpReserveId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeleteSslVpnServer
//
// PARAMS:
//   - request: the arguments to DeleteSslVpnServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSslVpnServer(request *DeleteSslVpnServerRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSslVpnServerUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.SslVpnServerId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeleteSslVpnUser
//
// PARAMS:
//   - request: the arguments to DeleteSslVpnUser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSslVpnUser(request *DeleteSslVpnUserRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSslVpnUserUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.UserId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeleteSubnet
//
// PARAMS:
//   - request: the arguments to DeleteSubnet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeleteUserGateway
//
// PARAMS:
//   - request: the arguments to DeleteUserGateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteUserGateway(request *DeleteUserGatewayRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteUserGatewayUri(VERSION_V1, util.StringValue(request.CgwId))).
		Do()
	return err
}

// DeleteVpc
//
// PARAMS:
//   - request: the arguments to DeleteVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteVpc(request *DeleteVpcRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeleteVpnTunnel
//
// PARAMS:
//   - request: the arguments to DeleteVpnTunnel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteVpnTunnel(request *DeleteVpnTunnelRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteVpnTunnelUri(VERSION_V1, util.StringValue(request.VpnConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
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
	return result, err
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
	return result, err
}

// ModifyGatewayLimitRules
//
// PARAMS:
//   - request: the arguments to ModifyGatewayLimitRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyGatewayLimitRules(request *ModifyGatewayLimitRulesRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyGatewayLimitRulesUri(VERSION_V1, util.StringValue(request.GlrId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// OpenVpcRelay
//
// PARAMS:
//   - request: the arguments to OpenVpcRelay
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) OpenVpcRelay(request *OpenVpcRelayRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getOpenVpcRelayUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
}

// ReleaseVpn
//
// PARAMS:
//   - request: the arguments to ReleaseVpn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseVpn(request *ReleaseVpnRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseVpnUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// RenewVpn
//
// PARAMS:
//   - request: the arguments to RenewVpn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenewVpn(request *RenewVpnRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenewVpnUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
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
	return result, err
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
	return result, err
}

// UnbindEip
//
// PARAMS:
//   - request: the arguments to UnbindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindEip(request *UnbindEipRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindEipUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// UpdateSslVpnServer
//
// PARAMS:
//   - request: the arguments to UpdateSslVpnServer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSslVpnServer(request *UpdateSslVpnServerRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSslVpnServerUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.SslVpnServerId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateSslVpnUsers
//
// PARAMS:
//   - request: the arguments to UpdateSslVpnUsers
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSslVpnUsers(request *UpdateSslVpnUsersRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSslVpnUsersUri(VERSION_V1, util.StringValue(request.VpnId), util.StringValue(request.UserId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateSubnet
//
// PARAMS:
//   - request: the arguments to UpdateSubnet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSubnet(request *UpdateSubnetRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateUserGateway
//
// PARAMS:
//   - request: the arguments to UpdateUserGateway
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateUserGateway(request *UpdateUserGatewayRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateUserGatewayUri(VERSION_V1, util.StringValue(request.CgwId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateVpc
//
// PARAMS:
//   - request: the arguments to UpdateVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpc(request *UpdateVpcRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateVpn
//
// PARAMS:
//   - request: the arguments to UpdateVpn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpn(request *UpdateVpnRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpnUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateVpnReleaseProtection
//
// PARAMS:
//   - request: the arguments to UpdateVpnReleaseProtection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpnReleaseProtection(request *UpdateVpnReleaseProtectionRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpnReleaseProtectionUri(VERSION_V1, util.StringValue(request.VpnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// UpdateVpnTunnel
//
// PARAMS:
//   - request: the arguments to UpdateVpnTunnel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateVpnTunnel(request *UpdateVpnTunnelRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpnTunnelUri(VERSION_V1, util.StringValue(request.VpnConnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
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
	return result, err
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
	return result, err
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
	return result, err
}
