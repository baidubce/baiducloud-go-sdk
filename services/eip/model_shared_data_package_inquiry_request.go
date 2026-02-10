package eip

type SharedDataPackageInquiryRequest struct {
	ClientToken       *string `json:"-"`
	ReservationLength *int32  `json:"reservationLength,omitempty"`
	Capacity          *string `json:"capacity,omitempty"`
	DeductPolicy      *string `json:"deductPolicy,omitempty"`
	PackageType       *string `json:"packageType,omitempty"`
}
