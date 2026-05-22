package cfw

type ListCfwRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
