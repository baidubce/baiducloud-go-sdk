package blb

type DeleteAppBlbServerGroupRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
	SgId        *string `json:"sgId,omitempty"`
}
