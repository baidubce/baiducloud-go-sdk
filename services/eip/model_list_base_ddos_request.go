package eip

type ListBaseDdosRequest struct {
	Ips     *string `json:"-"`
	Type    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
