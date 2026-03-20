package blb

type RefundBlbRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
