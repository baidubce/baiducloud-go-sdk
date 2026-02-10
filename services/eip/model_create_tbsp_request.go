package eip

type CreateTbspRequest struct {
	ClientToken         *string `json:"-"`
	Name                *string `json:"name,omitempty"`
	LineType            *string `json:"lineType,omitempty"`
	IpCapacity          *int32  `json:"ipCapacity,omitempty"`
	ReservationLength   *int32  `json:"reservationLength,omitempty"`
	ReservationTimeUnit *string `json:"reservationTimeUnit,omitempty"`
	AutoRenewTime       *int32  `json:"autoRenewTime,omitempty"`
	AutoRenewTimeUnit   *string `json:"autoRenewTimeUnit,omitempty"`
}
