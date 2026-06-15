package bcc

type GetInstanceNoChargeListRequest struct {
	Marker      *string `json:"-"`
	MaxKeys     *int32  `json:"-"`
	InternalIp  *string `json:"-"`
	KeypairId   *string `json:"-"`
	ZoneName    *string `json:"-"`
	InstanceIds *string `json:"-"`
}
