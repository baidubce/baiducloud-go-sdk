package snic

type Billing struct {
	PaymentTiming *string      `json:"paymentTiming,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}
