package snic

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// CreateSnic
//
// PARAMS:
//   - request: the arguments to CreateSnic
//
// RETURNS:
//   - CreateSnicResponse: The return type of the CreateSnic interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSnic(request *CreateSnicRequest) (*CreateSnicResponse, error) {
	result := &CreateSnicResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSnicUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSnic
//
// PARAMS:
//   - request: the arguments to DeleteSnic
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSnic(request *DeleteSnicRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSnicUri(VERSION_V1, util.StringValue(request.EndpointId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DescribeSnic
//
// PARAMS:
//   - request: the arguments to DescribeSnic
//
// RETURNS:
//   - DescribeSnicResponse: The return type of the DescribeSnic interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeSnic(request *DescribeSnicRequest) (*DescribeSnicResponse, error) {
	result := &DescribeSnicResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeSnicUri(VERSION_V1, util.StringValue(request.EndpointId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSnic
//
// PARAMS:
//   - request: the arguments to ListSnic
//
// RETURNS:
//   - ListSnicResponse: The return type of the ListSnic interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSnic(request *ListSnicRequest) (*ListSnicResponse, error) {
	result := &ListSnicResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSnicUri(VERSION_V1)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("ipAddress", util.StringValue(request.IpAddress)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("subnetId", util.StringValue(request.SubnetId)).
		WithQueryParamFilter("service", util.StringValue(request.Service)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAvailablePublicServices
//
// PARAMS:
//   - request: the arguments to QueryAvailablePublicServices
//
// RETURNS:
//   - QueryAvailablePublicServicesResponse: The return type of the QueryAvailablePublicServices interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAvailablePublicServices() (*QueryAvailablePublicServicesResponse, error) {
	result := &QueryAvailablePublicServicesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAvailablePublicServicesUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateSnic
//
// PARAMS:
//   - request: the arguments to UpdateSnic
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSnic(request *UpdateSnicRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSnicUri(VERSION_V1, util.StringValue(request.EndpointId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateSnicEsg
//
// PARAMS:
//   - request: the arguments to UpdateSnicEsg
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSnicEsg(request *UpdateSnicEsgRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSnicEsgUri(VERSION_V1, util.StringValue(request.EndpointId))).
		WithQueryParam("bindEsg", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateSnicSg
//
// PARAMS:
//   - request: the arguments to UpdateSnicSg
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSnicSg(request *UpdateSnicSgRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSnicSgUri(VERSION_V1, util.StringValue(request.EndpointId))).
		WithQueryParam("bindSg", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
