package blb

type UpdateAppBlbServerGroupRsRequest struct {
	BlbId             *string             `json:"-"`
	ClientToken       *string             `json:"-"`
	SgId              *string             `json:"sgId,omitempty"`
	BackendServerList []*AppBackendServer `json:"backendServerList,omitempty"`
}
