package rds

type QueryInstanceListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
	VnetIp  *string `json:"-"`
}
