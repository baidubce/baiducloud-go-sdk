package dns

type CreateZoneRequest struct {
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
}
