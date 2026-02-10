package eip

type AddTbspIpWhitelistRequest struct {
	Id          *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Ip          *string   `json:"ip,omitempty"`
	IpCidrList  []*string `json:"ipCidrList,omitempty"`
}
