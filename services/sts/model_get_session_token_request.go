package sts

type GetSessionTokenRequest struct {
	DurationSeconds   *string `json:"-"`
	AccessControlList *string `json:"accessControlList,omitempty"`
	Attachment        *string `json:"attachment,omitempty"`
}
