package rapidfs

type DescribeOrderRequest struct {
	Action  *string `json:"-"`
	OrderId *string `json:"orderId,omitempty"`
}
