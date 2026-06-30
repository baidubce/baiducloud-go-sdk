package iam

type DeleteAccessKeyRequest struct {
	UserName    *string `json:"-"`
	AccessKeyId *string `json:"-"`
}
