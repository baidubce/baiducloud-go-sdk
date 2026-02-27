package eip

type StopEipAutoRenewRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
