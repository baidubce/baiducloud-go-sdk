package csn

type BindCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Action      *string `json:"-"`
	CsnId       *string `json:"csnId,omitempty"`
}
