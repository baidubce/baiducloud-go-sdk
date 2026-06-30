package iam

type GetSessionApiKeyRequest struct {
	ExpireInSeconds *int32  `json:"-"`
	Acl             *string `json:"-"`
}
