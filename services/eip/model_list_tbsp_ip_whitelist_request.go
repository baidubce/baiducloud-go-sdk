package eip

type ListTbspIpWhitelistRequest struct {
	Id      *string `json:"-"`
	Ip      *string `json:"-"`
	IpCidr  *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
