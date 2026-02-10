package eip

type Billing struct {
	PaymentTiming *string      `json:"paymentTiming,omitempty"`
	BillingMethod *string      `json:"billingMethod,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}
