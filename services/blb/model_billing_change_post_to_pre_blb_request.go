package blb

type BillingChangePostToPreBlbRequest struct {
	BlbId             *string `json:"-"`
	ClientToken       *string `json:"-"`
	BillingMethod     *string `json:"billingMethod,omitempty"`
	PerformanceLevel  *string `json:"performanceLevel,omitempty"`
	ReservationLength *int32  `json:"reservationLength,omitempty"`
}
