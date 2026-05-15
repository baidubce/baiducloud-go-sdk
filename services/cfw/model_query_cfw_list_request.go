package cfw

type QueryCfwListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
