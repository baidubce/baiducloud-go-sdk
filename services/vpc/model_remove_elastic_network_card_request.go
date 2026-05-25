package vpc

type RemoveElasticNetworkCardRequest struct {
	EniId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
