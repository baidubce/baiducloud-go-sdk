package cfw

type CfwBind struct {
	Region     *string `json:"region,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	Role       *string `json:"role,omitempty"`
	MemberId   *string `json:"memberId,omitempty"`
}
