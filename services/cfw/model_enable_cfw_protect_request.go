package cfw

type EnableCfwProtectRequest struct {
	CfwId      *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	Role       *string `json:"role,omitempty"`
	MemberId   *string `json:"memberId,omitempty"`
}
