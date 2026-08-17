package aigw

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

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
		WithBody(request).
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
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
