package sts

type AssumeRoleRequest struct {
	DurationSeconds   *string `json:"-"`
	AccountId         *string `json:"-"`
	UserId            *string `json:"-"`
	RoleName          *string `json:"-"`
	AccessControlList *string `json:"accessControlList,omitempty"`
}
