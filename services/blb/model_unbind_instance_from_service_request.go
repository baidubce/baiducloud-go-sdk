package blb

type UnbindInstanceFromServiceRequest struct {
	Service     *string `json:"-"`
	ClientToken *string `json:"-"`
}
