package scs

type InstanceListRequest struct {
	Marker      *string `json:"-"`
	MaxKeys     *string `json:"-"`
	InstanceIds *string `json:"-"`
	VnetIp      *string `json:"-"`
}
