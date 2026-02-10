package eip

type Reservation struct {
	ReservationLength   *int32  `json:"reservationLength,omitempty"`
	ReservationTimeUnit *string `json:"reservationTimeUnit,omitempty"`
}
