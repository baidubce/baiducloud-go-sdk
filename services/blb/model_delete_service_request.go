package blb

type DeleteServiceRequest struct {
	Service     *string `json:"-"`
	ClientToken *string `json:"-"`
}
