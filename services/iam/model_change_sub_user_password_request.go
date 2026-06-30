package iam

type ChangeSubUserPasswordRequest struct {
	UserName *string `json:"-"`
	Password *string `json:"password,omitempty"`
}
