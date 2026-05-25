package vpc

type DeleteElasticNetworkCardAuxiliaryIpRequest struct {
	EniId            *string `json:"-"`
	PrivateIpAddress *string `json:"-"`
	ClientToken      *string `json:"-"`
}
