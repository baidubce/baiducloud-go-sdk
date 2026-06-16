package cfw

type ListStatelessCfwRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
