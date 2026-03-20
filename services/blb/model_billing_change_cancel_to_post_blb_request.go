package blb

type BillingChangeCancelToPostBlbRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
