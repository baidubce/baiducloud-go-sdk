package blb

type BindInstanceToServiceRequest struct {
	Service     *string `json:"-"`
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
}
