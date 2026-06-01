package blb

type UpdateServiceAuthRequest struct {
	Service     *string `json:"-"`
	ClientToken *string `json:"-"`
	AuthList    []*Auth `json:"authList,omitempty"`
}
