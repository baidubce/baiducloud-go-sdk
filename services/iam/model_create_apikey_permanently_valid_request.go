package iam

type CreateApikeyPermanentlyValidRequest struct {
	UserId *string `json:"userId,omitempty"`
	Acl    *string `json:"acl,omitempty"`
	Name   *string `json:"name,omitempty"`
}
