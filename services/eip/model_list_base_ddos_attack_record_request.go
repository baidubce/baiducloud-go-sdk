package eip

type ListBaseDdosAttackRecordRequest struct {
	Ip        *string `json:"-"`
	StartTime *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
}
