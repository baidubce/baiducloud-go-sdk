package blb

type DeleteAppBlbServerGroupPortRequest struct {
	BlbId       *string   `json:"-"`
	ClientToken *string   `json:"-"`
	SgId        *string   `json:"sgId,omitempty"`
	PortIdList  []*string `json:"portIdList,omitempty"`
}
