package blb

type UpdateAppBlbServerGroupRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
	SgId        *string `json:"sgId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Desc        *string `json:"desc,omitempty"`
}
