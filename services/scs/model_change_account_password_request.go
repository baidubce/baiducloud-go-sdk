package scs

type ChangeAccountPasswordRequest struct {
	InstanceId *string `json:"-"`
	UserName   *string `json:"userName,omitempty"`
	ClientAuth *string `json:"clientAuth,omitempty"`
}
