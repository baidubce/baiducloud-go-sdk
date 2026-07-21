package iam

type CreateApikeyPermanentlyValidRequest struct {
	UserId *string `json:"userId,omitempty"`
	Acl    *ACL    `json:"acl,omitempty"`
	Name   *string `json:"name,omitempty"`
}
