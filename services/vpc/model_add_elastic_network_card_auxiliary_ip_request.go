package vpc

type AddElasticNetworkCardAuxiliaryIpRequest struct {
	EniId            *string `json:"-"`
	ClientToken      *string `json:"-"`
	IsIpv6           *bool   `json:"isIpv6,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
}
