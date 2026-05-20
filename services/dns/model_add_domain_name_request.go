package dns

type AddDomainNameRequest struct {
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
}
