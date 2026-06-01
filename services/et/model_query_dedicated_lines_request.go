package et

type QueryDedicatedLinesRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
	Status  *string `json:"-"`
}
