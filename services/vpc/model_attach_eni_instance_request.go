package vpc

type AttachEniInstanceRequest struct {
	EniId        *string `json:"-"`
	ClientToken  *string `json:"-"`
	InstanceId   *string `json:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
}
