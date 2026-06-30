package iam

type AccountCountInfo struct {
	UserCount   *int32 `json:"userCount,omitempty"`
	PolicyCount *int32 `json:"policyCount,omitempty"`
	GroupCount  *int32 `json:"groupCount,omitempty"`
}
