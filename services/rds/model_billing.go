package rds

type Billing struct {
	PaymentTiming *string      `json:"paymentTiming,omitempty"`
	Reservation   *interface{} `json:"reservation,omitempty"`
}
