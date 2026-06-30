package iam

type UnbindSubUserVirtualMfaRequest struct {
	UserName *string `json:"-"`
	MfaType  *string `json:"-"`
}
