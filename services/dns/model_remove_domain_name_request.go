package dns

type RemoveDomainNameRequest struct {
	ZoneName    *string `json:"-"`
	ClientToken *string `json:"-"`
}
