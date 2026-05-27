package privatezone

type CreatePrivateZoneRequest struct {
	ClientToken *string `json:"-"`
	ZoneName    *string `json:"zoneName,omitempty"`
}
