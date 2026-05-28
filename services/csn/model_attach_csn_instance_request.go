package csn

type AttachCsnInstanceRequest struct {
	CsnId             *string `json:"-"`
	ClientToken       *string `json:"-"`
	InstanceType      *string `json:"instanceType,omitempty"`
	InstanceId        *string `json:"instanceId,omitempty"`
	InstanceRegion    *string `json:"instanceRegion,omitempty"`
	InstanceAccountId *string `json:"instanceAccountId,omitempty"`
}
