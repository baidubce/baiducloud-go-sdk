package privatezone

type CreateAPrivateZoneRequest struct {
	ClientToken *string `json:"-"`
	ZoneName    *string `json:"zoneName,omitempty"`
}
