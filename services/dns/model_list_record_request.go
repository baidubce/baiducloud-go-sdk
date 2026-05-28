package dns

type ListRecordRequest struct {
	ZoneName *string `json:"-"`
	Rr       *string `json:"-"`
	Id       *string `json:"-"`
	Marker   *string `json:"-"`
	MaxKeys  *int32  `json:"-"`
}
