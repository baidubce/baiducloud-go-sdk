package bls

type ListLogShipperRequest struct {
	LogShipperID   *string `json:"-"`
	LogShipperName *string `json:"-"`
	Project        *string `json:"-"`
	LogStoreName   *string `json:"-"`
	DestType       *string `json:"-"`
	Status         *string `json:"-"`
	OrderBy        *string `json:"-"`
	Order          *string `json:"-"`
	PageNo         *int32  `json:"-"`
	PageSize       *int32  `json:"-"`
}
