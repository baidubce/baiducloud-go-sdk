

package vpc




type Billing struct {
	PaymentTiming *string `json:"paymentTiming,omitempty"`
	Reservation *Reservation `json:"reservation,omitempty"`
}

