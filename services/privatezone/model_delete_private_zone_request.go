package privatezone

type DeletePrivateZoneRequest struct {
	ZoneId      *string `json:"-"`
	ClientToken *string `json:"-"`
}
