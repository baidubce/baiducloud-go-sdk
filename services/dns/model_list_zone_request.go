package dns

type ListZoneRequest struct {
	Name    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
