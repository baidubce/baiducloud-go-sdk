package scs

type CreateAccountRequest struct {
	InstanceId *string `json:"-"`
	UserName   *string `json:"userName,omitempty"`
	ClientAuth *string `json:"clientAuth,omitempty"`
	Extra      *string `json:"extra,omitempty"`
	UserType   *int32  `json:"userType,omitempty"`
}
