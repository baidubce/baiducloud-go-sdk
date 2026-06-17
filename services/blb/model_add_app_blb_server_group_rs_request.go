package blb

type AddAppBlbServerGroupRsRequest struct {
	BlbId             *string                      `json:"-"`
	ClientToken       *string                      `json:"-"`
	SgId              *string                      `json:"sgId,omitempty"`
	BackendServerList []*AppBackendServerForCreate `json:"backendServerList,omitempty"`
}
