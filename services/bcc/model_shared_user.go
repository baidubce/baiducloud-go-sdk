package bcc

type SharedUser struct {
	Account   *string `json:"account,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	UcAccount *string `json:"ucAccount,omitempty"`
}
