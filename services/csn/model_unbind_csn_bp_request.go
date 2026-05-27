package csn

type UnbindCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Action      *string `json:"-"`
	CsnId       *string `json:"csnId,omitempty"`
}
