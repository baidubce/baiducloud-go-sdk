package vpc

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"strings"
)

const (
	VERSION_V1 = "v1"
)

// CloseVpcRelay
//
// PARAMS:
//   - request: the arguments to CloseVpcRelay
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CloseVpcRelay(request *CloseVpcRelayRequest) error {
	if request.VpcId == nil {
		return fmt.Errorf("vpcId is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCloseVpcRelayUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// CreateAReservedNetworkSegment
//
// PARAMS:
//   - request: the arguments to CreateAReservedNetworkSegment
//
// RETURNS:
//   - CreateAReservedNetworkSegmentResponse: The return type of the CreateAReservedNetworkSegment interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAReservedNetworkSegment(request *CreateAReservedNetworkSegmentRequest) (*CreateAReservedNetworkSegmentResponse, error) {
	if request.SubnetId == nil {
		return nil, fmt.Errorf("SubnetId is required and must be specified")
	}
	if request.IpCidr == nil {
		return nil, fmt.Errorf("IpCidr is required and must be specified")
	}
	if request.IpVersion == nil {
		return nil, fmt.Errorf("IpVersion is required and must be specified")
	}
	result := &CreateAReservedNetworkSegmentResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAReservedNetworkSegmentUri(VERSION_V1)).
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
	if request.Name == nil {
		return nil, fmt.Errorf("Name is required and must be specified")
	}
	if request.ZoneName == nil {
		return nil, fmt.Errorf("ZoneName is required and must be specified")
	}
	if request.Cidr == nil {
		return nil, fmt.Errorf("Cidr is required and must be specified")
	}
	if request.VpcId == nil {
		return nil, fmt.Errorf("VpcId is required and must be specified")
	}
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

// CreateVpc
//
// PARAMS:
//   - request: the arguments to CreateVpc
//
// RETURNS:
//   - CreateVpcResponse: The return type of the CreateVpc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVpc(request *CreateVpcRequest) (*CreateVpcResponse, error) {
	if request.Name == nil {
		return nil, fmt.Errorf("Name is required and must be specified")
	}
	if request.Cidr == nil {
		return nil, fmt.Errorf("Cidr is required and must be specified")
	}
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

// DeleteReservedNetworkSegment
//
// PARAMS:
//   - request: the arguments to DeleteReservedNetworkSegment
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteReservedNetworkSegment(request *DeleteReservedNetworkSegmentRequest) error {
	if request.IpReserveId == nil {
		return fmt.Errorf("ipReserveId is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteReservedNetworkSegmentUri(VERSION_V1, util.StringValue(request.IpReserveId))).
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
	if request.SubnetId == nil {
		return fmt.Errorf("subnetId is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
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
	if request.VpcId == nil {
		return fmt.Errorf("vpcId is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// EnableVpcRelay
//
// PARAMS:
//   - request: the arguments to EnableVpcRelay
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableVpcRelay(request *EnableVpcRelayRequest) error {
	if request.VpcId == nil {
		return fmt.Errorf("vpcId is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableVpcRelayUri(VERSION_V1, util.StringValue(request.VpcId))).
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
	if request.SubnetId == nil {
		return nil, fmt.Errorf("subnetId is required and must be specified")
	}
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
	if request.VpcId == nil {
		return nil, fmt.Errorf("vpcId is required and must be specified")
	}
	result := &QuerySpecifiedVpcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySpecifiedVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
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

// QueryTheIpAddressesOccupiedByProductsWithinVpc
//
// PARAMS:
//   - request: the arguments to QueryTheIpAddressesOccupiedByProductsWithinVpc
//
// RETURNS:
//   - QueryTheIpAddressesOccupiedByProductsWithinVpcResponse: The return type of the QueryTheIpAddressesOccupiedByProductsWithinVpc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheIpAddressesOccupiedByProductsWithinVpc(request *QueryTheIpAddressesOccupiedByProductsWithinVpcRequest) (*QueryTheIpAddressesOccupiedByProductsWithinVpcResponse, error) {
	if request.VpcId == nil {
		return nil, fmt.Errorf("vpcId is required and must be specified")
	}
	result := &QueryTheIpAddressesOccupiedByProductsWithinVpcResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheIpAddressesOccupiedByProductsWithinVpcUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("subnetId", util.StringValue(request.SubnetId)).
		WithQueryParamFilter("resourceType", util.StringValue(request.ResourceType)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

// QueryTheReservedNetworkSegmentList
//
// PARAMS:
//   - request: the arguments to QueryTheReservedNetworkSegmentList
//
// RETURNS:
//   - QueryTheReservedNetworkSegmentListResponse: The return type of the QueryTheReservedNetworkSegmentList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheReservedNetworkSegmentList(request *QueryTheReservedNetworkSegmentListRequest) (*QueryTheReservedNetworkSegmentListResponse, error) {
	result := &QueryTheReservedNetworkSegmentListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheReservedNetworkSegmentListUri(VERSION_V1)).
		WithQueryParamFilter("subnetId", util.StringValue(request.SubnetId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
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
	if request.VpcId == nil {
		return nil, fmt.Errorf("vpcId is required and must be specified")
	}
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

// UpdateSubnet
//
// PARAMS:
//   - request: the arguments to UpdateSubnet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSubnet(request *UpdateSubnetRequest) error {
	if request.SubnetId == nil {
		return fmt.Errorf("subnetId is required and must be specified")
	}
	if request.Name == nil {
		return fmt.Errorf("Name is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSubnetUri(VERSION_V1, util.StringValue(request.SubnetId))).
		WithQueryParam("modifyAttribute", "").
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
	if request.VpcId == nil {
		return fmt.Errorf("vpcId is required and must be specified")
	}
	if request.Name == nil {
		return fmt.Errorf("Name is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateVpcUri(VERSION_V1, util.StringValue(request.VpcId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}
