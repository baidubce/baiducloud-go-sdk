package bls

type ListLogShipperRecordRequest struct {
	LogShipperID *string `json:"-"`
	SinceHours   *int32  `json:"-"`
	PageNo       *int32  `json:"-"`
	PageSize     *int32  `json:"-"`
}
