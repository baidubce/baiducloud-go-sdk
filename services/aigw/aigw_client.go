package aigw

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// CreateAIGateway
//
// PARAMS:
//   - request: the arguments to CreateAIGateway
//
// RETURNS:
//   - CreateAIGatewayResponse: The return type of the CreateAIGateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAIGateway(request *CreateAIGatewayRequest) (*CreateAIGatewayResponse, error) {
	result := &CreateAIGatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAIGatewayUri()).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateConsumer
//
// PARAMS:
//   - request: the arguments to CreateConsumer
//
// RETURNS:
//   - CreateConsumerResponse: The return type of the CreateConsumer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (*CreateConsumerResponse, error) {
	result := &CreateConsumerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateConsumerUri(util.StringValue(request.InstanceId))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRoute
//
// PARAMS:
//   - request: the arguments to CreateRoute
//
// RETURNS:
//   - CreateRouteResponse: The return type of the CreateRoute interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateRoute(request *CreateRouteRequest) (*CreateRouteResponse, error) {
	result := &CreateRouteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRouteUri(util.StringValue(request.InstanceId), util.StringValue(request.ClusterId))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateService
//
// PARAMS:
//   - request: the arguments to CreateService
//
// RETURNS:
//   - CreateServiceResponse: The return type of the CreateService interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateService(request *CreateServiceRequest) (*CreateServiceResponse, error) {
	result := &CreateServiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateServiceUri(util.StringValue(request.InstanceId))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAIGateway
//
// PARAMS:
//   - request: the arguments to DeleteAIGateway
//
// RETURNS:
//   - DeleteAIGatewayResponse: The return type of the DeleteAIGateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAIGateway(request *DeleteAIGatewayRequest) (*DeleteAIGatewayResponse, error) {
	result := &DeleteAIGatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAIGatewayUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("force", util.BoolValue(request.Force)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteConsumer
//
// PARAMS:
//   - request: the arguments to DeleteConsumer
//
// RETURNS:
//   - DeleteConsumerResponse: The return type of the DeleteConsumer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteConsumer(request *DeleteConsumerRequest) (*DeleteConsumerResponse, error) {
	result := &DeleteConsumerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteConsumerUri(util.StringValue(request.InstanceId), util.StringValue(request.ConsumerId))).
		WithQueryParamFilter("keyType", util.StringValue(request.KeyType)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteRoute
//
// PARAMS:
//   - request: the arguments to DeleteRoute
//
// RETURNS:
//   - DeleteRouteResponse: The return type of the DeleteRoute interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteRoute(request *DeleteRouteRequest) (*DeleteRouteResponse, error) {
	result := &DeleteRouteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRouteUri(util.StringValue(request.InstanceId), util.StringValue(request.RouteName))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteService
//
// PARAMS:
//   - request: the arguments to DeleteService
//
// RETURNS:
//   - DeleteServiceResponse: The return type of the DeleteService interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteService(request *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	result := &DeleteServiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteServiceUri(util.StringValue(request.InstanceId), util.StringValue(request.ServiceName), util.StringValue(request.Namespace))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAIGatewayDetail
//
// PARAMS:
//   - request: the arguments to GetAIGatewayDetail
//
// RETURNS:
//   - GetAIGatewayDetailResponse: The return type of the GetAIGatewayDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAIGatewayDetail(request *GetAIGatewayDetailRequest) (*GetAIGatewayDetailResponse, error) {
	result := &GetAIGatewayDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAIGatewayDetailUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("srcProduct", util.StringValue(request.SrcProduct)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConsumer
//
// PARAMS:
//   - request: the arguments to GetConsumer
//
// RETURNS:
//   - GetConsumerResponse: The return type of the GetConsumer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetConsumer(request *GetConsumerRequest) (*GetConsumerResponse, error) {
	result := &GetConsumerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetConsumerUri(util.StringValue(request.InstanceId), util.StringValue(request.ConsumerId))).
		WithQueryParamFilter("keyType", util.StringValue(request.KeyType)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConsumerList
//
// PARAMS:
//   - request: the arguments to GetConsumerList
//
// RETURNS:
//   - GetConsumerListResponse: The return type of the GetConsumerList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetConsumerList(request *GetConsumerListRequest) (*GetConsumerListResponse, error) {
	result := &GetConsumerListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetConsumerListUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("tagKey", util.StringValue(request.TagKey)).
		WithQueryParamFilter("tagValue", util.StringValue(request.TagValue)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetServiceDetail
//
// PARAMS:
//   - request: the arguments to GetServiceDetail
//
// RETURNS:
//   - GetServiceDetailResponse: The return type of the GetServiceDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetServiceDetail(request *GetServiceDetailRequest) (*GetServiceDetailResponse, error) {
	result := &GetServiceDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetServiceDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.ServiceName))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetServiceList
//
// PARAMS:
//   - request: the arguments to GetServiceList
//
// RETURNS:
//   - GetServiceListResponse: The return type of the GetServiceList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetServiceList(request *GetServiceListRequest) (*GetServiceListResponse, error) {
	result := &GetServiceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetServiceListUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("serviceSource", util.StringValue(request.ServiceSource)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAIGateways
//
// PARAMS:
//   - request: the arguments to ListAIGateways
//
// RETURNS:
//   - ListAIGatewaysResponse: The return type of the ListAIGateways interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAIGateways(request *ListAIGatewaysRequest) (*ListAIGatewaysResponse, error) {
	result := &ListAIGatewaysResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAIGatewaysUri()).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("srcProduct", util.StringValue(request.SrcProduct)).
		WithQueryParamFilter("tagKey", util.StringValue(request.TagKey)).
		WithQueryParamFilter("tagValue", util.StringValue(request.TagValue)).
		WithQueryParamFilter("resourceGroupId", util.StringValue(request.ResourceGroupId)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListServicesBySource
//
// PARAMS:
//   - request: the arguments to ListServicesBySource
//
// RETURNS:
//   - ListServicesBySourceResponse: The return type of the ListServicesBySource interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListServicesBySource(request *ListServicesBySourceRequest) (*ListServicesBySourceResponse, error) {
	result := &ListServicesBySourceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListServicesBySourceUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("serviceSource", util.StringValue(request.ServiceSource)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRoutingDetails
//
// PARAMS:
//   - request: the arguments to QueryRoutingDetails
//
// RETURNS:
//   - QueryRoutingDetailsResponse: The return type of the QueryRoutingDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRoutingDetails(request *QueryRoutingDetailsRequest) (*QueryRoutingDetailsResponse, error) {
	result := &QueryRoutingDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRoutingDetailsUri(util.StringValue(request.InstanceId), util.StringValue(request.RouteName))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRoutingList
//
// PARAMS:
//   - request: the arguments to QueryRoutingList
//
// RETURNS:
//   - QueryRoutingListResponse: The return type of the QueryRoutingList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRoutingList(request *QueryRoutingListRequest) (*QueryRoutingListResponse, error) {
	result := &QueryRoutingListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRoutingListUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("routeName", util.StringValue(request.RouteName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAIGateway
//
// PARAMS:
//   - request: the arguments to UpdateAIGateway
//
// RETURNS:
//   - UpdateAIGatewayResponse: The return type of the UpdateAIGateway interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAIGateway(request *UpdateAIGatewayRequest) (*UpdateAIGatewayResponse, error) {
	result := &UpdateAIGatewayResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAIGatewayUri(util.StringValue(request.InstanceId))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateConsumer
//
// PARAMS:
//   - request: the arguments to UpdateConsumer
//
// RETURNS:
//   - UpdateConsumerResponse: The return type of the UpdateConsumer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateConsumer(request *UpdateConsumerRequest) (*UpdateConsumerResponse, error) {
	result := &UpdateConsumerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateConsumerUri(util.StringValue(request.InstanceId), util.StringValue(request.ConsumerId))).
		WithQueryParamFilter("keyType", util.StringValue(request.KeyType)).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateRoute
//
// PARAMS:
//   - request: the arguments to UpdateRoute
//
// RETURNS:
//   - UpdateRouteResponse: The return type of the UpdateRoute interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateRoute(request *UpdateRouteRequest) (*UpdateRouteResponse, error) {
	result := &UpdateRouteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRouteUri(util.StringValue(request.InstanceId), util.StringValue(request.RouteName))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateService
//
// PARAMS:
//   - request: the arguments to UpdateService
//
// RETURNS:
//   - UpdateServiceResponse: The return type of the UpdateService interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateService(request *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	result := &UpdateServiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateServiceUri(util.StringValue(request.InstanceId), util.StringValue(request.ServiceNamePath))).
		WithHeaderFilter("X-Region", util.StringValue(request.XRegion)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
