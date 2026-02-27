package eip

type DirectEipRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
