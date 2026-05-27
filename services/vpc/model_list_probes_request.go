package vpc

type ListProbesRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
