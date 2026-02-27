package eip

type UnDirectEipRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
