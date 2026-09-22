package scs

type DeleteAccountRequest struct {
	InstanceId *string `json:"-"`
	UserName   *string `json:"userName,omitempty"`
}
