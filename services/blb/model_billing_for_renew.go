package blb

type BillingForRenew struct {
	Reservation *ReservationForCreate `json:"reservation,omitempty"`
}
