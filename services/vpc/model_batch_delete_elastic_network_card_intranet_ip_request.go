package vpc

type BatchDeleteElasticNetworkCardIntranetIpRequest struct {
	EniId              *string   `json:"-"`
	ClientToken        *string   `json:"-"`
	PrivateIpAddresses []*string `json:"privateIpAddresses,omitempty"`
}
