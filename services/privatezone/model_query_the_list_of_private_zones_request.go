package privatezone

type QueryTheListOfPrivateZonesRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
