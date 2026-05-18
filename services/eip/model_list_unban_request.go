package eip

type ListUnbanRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
	Ip      *string `json:"-"`
}
