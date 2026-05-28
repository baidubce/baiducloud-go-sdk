package csn

type UnbindCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	CsnId       *string `json:"csnId,omitempty"`
}
