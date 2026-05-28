package dns

type ListLineGroupRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
