package csn

type BindCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	CsnId       *string `json:"csnId,omitempty"`
}
