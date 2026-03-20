package vpc

type SslVpnUserVo struct {
	ClientName  *string `json:"clientName,omitempty"`
	UserName    *string `json:"userName,omitempty"`
	UserId      *string `json:"userId,omitempty"`
	Description *string `json:"description,omitempty"`
}
