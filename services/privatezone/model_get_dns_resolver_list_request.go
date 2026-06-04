package privatezone

type GetDnsResolverListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *string `json:"-"`
	Status  *string `json:"-"`
}
