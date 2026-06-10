package bcc

type ListVolumesRequest struct {
	Marker          *string `json:"-"`
	MaxKeys         *int32  `json:"-"`
	InstanceId      *string `json:"-"`
	ZoneName        *string `json:"-"`
	ClusterId       *string `json:"-"`
	ChargeFilter    *string `json:"-"`
	UsageFilter     *string `json:"-"`
	Name            *string `json:"-"`
	ProductCategory *string `json:"-"`
}
