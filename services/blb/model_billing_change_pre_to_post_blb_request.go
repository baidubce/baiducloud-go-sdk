package blb

type BillingChangePreToPostBlbRequest struct {
	BlbId                *string `json:"-"`
	ClientToken          *string `json:"-"`
	BillingMethod        *string `json:"billingMethod,omitempty"`
	PerformanceLevel     *string `json:"performanceLevel,omitempty"`
	EffectiveImmediately *bool   `json:"effectiveImmediately,omitempty"`
}
