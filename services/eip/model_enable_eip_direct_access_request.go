package eip

type EnableEipDirectAccessRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
