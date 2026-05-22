package blb

type AppPolicyForUpdate struct {
	PolicyId    *string `json:"policyId,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
	Description *string `json:"description,omitempty"`
}
