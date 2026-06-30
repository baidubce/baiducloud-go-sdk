package iam

type AccountLimitInfo struct {
	UserLimit                      *int32 `json:"userLimit,omitempty"`
	PolicyLimit                    *int32 `json:"policyLimit,omitempty"`
	ContactsLimit                  *int32 `json:"contactsLimit,omitempty"`
	GroupLimit                     *int32 `json:"groupLimit,omitempty"`
	SubUserOfGroupLimit            *int32 `json:"subUserOfGroupLimit,omitempty"`
	GroupMaxAttachPolicyLimit      *int32 `json:"groupMaxAttachPolicyLimit,omitempty"`
	UserRolePerAccountLimit        *int32 `json:"userRolePerAccountLimit,omitempty"`
	RoleMaxAttachSystemPolicyLimit *int32 `json:"roleMaxAttachSystemPolicyLimit,omitempty"`
	RoleMaxAttachCustomPolicyLimit *int32 `json:"roleMaxAttachCustomPolicyLimit,omitempty"`
	AkskLimit                      *int32 `json:"akskLimit,omitempty"`
}
