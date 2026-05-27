package csn

type DetachCsnInstanceRequest struct {
	CsnId             *string `json:"-"`
	ClientToken       *string `json:"-"`
	Action            *string `json:"-"`
	InstanceType      *string `json:"instanceType,omitempty"`
	InstanceId        *string `json:"instanceId,omitempty"`
	InstanceRegion    *string `json:"instanceRegion,omitempty"`
	InstanceAccountId *string `json:"instanceAccountId,omitempty"`
}
