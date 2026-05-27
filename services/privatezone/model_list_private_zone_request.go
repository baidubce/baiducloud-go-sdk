package privatezone

type ListPrivateZoneRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
