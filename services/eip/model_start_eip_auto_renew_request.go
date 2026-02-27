package eip

type StartEipAutoRenewRequest struct {
	Eip               *string `json:"-"`
	ClientToken       *string `json:"-"`
	AutoRenewTimeUnit *string `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     *int32  `json:"autoRenewTime,omitempty"`
}
