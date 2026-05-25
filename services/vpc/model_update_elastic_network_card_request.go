package vpc

type UpdateElasticNetworkCardRequest struct {
	EniId       *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
