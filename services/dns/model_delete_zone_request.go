package dns

type DeleteZoneRequest struct {
	ZoneName    *string `json:"-"`
	ClientToken *string `json:"-"`
}
