package eip

type RemoveTbspIpWhitelistRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	Ip          *string `json:"-"`
	WhitelistId *string `json:"-"`
}
