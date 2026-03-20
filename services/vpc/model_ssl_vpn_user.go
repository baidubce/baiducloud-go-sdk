package vpc

type SslVpnUser struct {
	UserName    *string `json:"userName,omitempty"`
	Password    *string `json:"password,omitempty"`
	Description *string `json:"description,omitempty"`
}
