package aigw

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "aigw." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V1 = "v1"

	CONSTANT_AIGW = "aigw"

	CONSTANT_ROUTE = "route"

	CONSTANT_DETAIL = "detail"

	CONSTANT_CONSUMER = "consumer"

	CONSTANT_CLUSTER = "cluster"

	CONSTANT_CONSUMERS = "consumers"
)

// Client of aigw service is a kind of BceClient, so derived from BceClient
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

func getCreateConsumerUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CONSUMER
}
func getCreateRouteUri(InstanceId string, ClusterId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + ClusterId + bce.URI_PREFIX + CONSTANT_ROUTE
}
func getDeleteConsumerUri(InstanceId string, ConsumerId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CONSUMER + bce.URI_PREFIX + ConsumerId
}
func getDeleteRouteUri(InstanceId string, RouteName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + RouteName + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_DETAIL
}
func getGetConsumerUri(InstanceId string, ConsumerId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CONSUMER + bce.URI_PREFIX + ConsumerId
}
func getGetConsumerListUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CONSUMERS
}
func getQueryRoutingDetailsUri(InstanceId string, RouteName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + RouteName + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_DETAIL
}
func getQueryRoutingListUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ROUTE
}
func getUpdateConsumerUri(InstanceId string, ConsumerId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CONSUMER + bce.URI_PREFIX + ConsumerId
}
func getUpdateRouteUri(InstanceId string, RouteName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIGW + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + RouteName + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_DETAIL
}
