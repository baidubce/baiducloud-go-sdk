package iam

type DisableAccessKeyRequest struct {
	UserName    *string `json:"-"`
	AccessKeyId *string `json:"-"`
}
