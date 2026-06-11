package bcc

type UnShareImageRequest struct {
	ImageId   *string `json:"-"`
	Account   *string `json:"account,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	UcAccount *string `json:"ucAccount,omitempty"`
}
