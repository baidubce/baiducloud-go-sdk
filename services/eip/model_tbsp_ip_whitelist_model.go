package eip

type TbspIpWhitelistModel struct {
	Ip          *string `json:"ip,omitempty"`
	WhitelistId *string `json:"whitelistId,omitempty"`
	IpCidr      *string `json:"ipCidr,omitempty"`
}
