package blb

type BillingForCreate struct {
	PaymentTiming *string               `json:"paymentTiming,omitempty"`
	BillingMethod *string               `json:"billingMethod,omitempty"`
	Reservation   *ReservationForCreate `json:"reservation,omitempty"`
}
