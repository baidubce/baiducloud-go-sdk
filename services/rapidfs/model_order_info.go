package rapidfs

type OrderInfo struct {
	OrderId     *string `json:"orderId,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty"`
	OrderStatus *string `json:"orderStatus,omitempty"`
}
