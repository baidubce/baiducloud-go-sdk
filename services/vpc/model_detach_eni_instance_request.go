package vpc

type DetachEniInstanceRequest struct {
	EniId       *string `json:"-"`
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
}
