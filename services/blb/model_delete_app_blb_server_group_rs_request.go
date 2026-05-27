package blb

type DeleteAppBlbServerGroupRsRequest struct {
	BlbId               *string   `json:"-"`
	ClientToken         *string   `json:"-"`
	SgId                *string   `json:"sgId,omitempty"`
	BackendServerIdList []*string `json:"backendServerIdList,omitempty"`
}
