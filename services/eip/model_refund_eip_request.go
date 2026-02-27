package eip

type RefundEipRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
