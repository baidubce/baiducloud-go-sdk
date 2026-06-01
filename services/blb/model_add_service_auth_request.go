package blb

type AddServiceAuthRequest struct {
	Service     *string `json:"-"`
	ClientToken *string `json:"-"`
	AuthList    []*Auth `json:"authList,omitempty"`
}
