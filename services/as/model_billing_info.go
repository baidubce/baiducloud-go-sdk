package as

type BillingInfo struct {
	PaymentTiming *string      `json:"paymentTiming,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}
