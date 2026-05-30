package pfs

type InstanceModel struct {
	Capacity       *int32        `json:"capacity,omitempty"`
	CreateTime     *string       `json:"createTime,omitempty"`
	Description    *string       `json:"description,omitempty"`
	Endpoint       *string       `json:"endpoint,omitempty"`
	InstanceId     *string       `json:"instanceId,omitempty"`
	InstanceStatus *string       `json:"instanceStatus,omitempty"`
	InstanceType   *string       `json:"instanceType,omitempty"`
	Name           *string       `json:"name,omitempty"`
	PaymentTiming  *string       `json:"paymentTiming,omitempty"`
	SubnetModel    *SubnetDetail `json:"subnetModel,omitempty"`
	Usage          *int32        `json:"usage,omitempty"`
	VpcId          *string       `json:"vpcId,omitempty"`
}
