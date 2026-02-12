package eip

type BindEipRequest struct {
	Eip          *string `json:"-"`
	ClientToken  *string `json:"-"`
	InstanceType *string `json:"instanceType,omitempty"`
	InstanceId   *string `json:"instanceId,omitempty"`
	InstanceIp   *string `json:"instanceIp,omitempty"`
}
