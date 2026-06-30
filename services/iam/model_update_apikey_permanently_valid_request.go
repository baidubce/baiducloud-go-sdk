package iam

type UpdateApikeyPermanentlyValidRequest struct {
	UserId *string `json:"userId,omitempty"`
	Id     *string `json:"id,omitempty"`
	Acl    *string `json:"acl,omitempty"`
}
