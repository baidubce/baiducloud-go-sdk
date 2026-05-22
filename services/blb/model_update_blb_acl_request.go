package blb

type UpdateBlbAclRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
	SupportAcl  *bool   `json:"supportAcl,omitempty"`
}
