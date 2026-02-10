package eip

type ListTbspIpCleanRequest struct {
	Id      *string `json:"-"`
	Ip      *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
