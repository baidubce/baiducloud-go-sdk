package iam

type EnableAccessKeyRequest struct {
	UserName    *string `json:"-"`
	AccessKeyId *string `json:"-"`
}
