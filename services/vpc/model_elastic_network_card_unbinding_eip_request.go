package vpc

type ElasticNetworkCardUnbindingEipRequest struct {
	EniId           *string `json:"-"`
	ClientToken     *string `json:"-"`
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
}
