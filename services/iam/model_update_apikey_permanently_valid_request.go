package iam

type UpdateApikeyPermanentlyValidRequest struct {
	UserId *string `json:"userId,omitempty"`
	Id     *string `json:"id,omitempty"`
	Acl    *ACL    `json:"acl,omitempty"`
}
