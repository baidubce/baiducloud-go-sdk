package vpc

type BatchIncreaseElasticNetworkCardIntranetIpRequest struct {
	EniId                 *string   `json:"-"`
	ClientToken           *string   `json:"-"`
	IsIpv6                *bool     `json:"isIpv6,omitempty"`
	PrivateIpAddresses    []*string `json:"privateIpAddresses,omitempty"`
	PrivateIpAddressCount *int32    `json:"privateIpAddressCount,omitempty"`
}
