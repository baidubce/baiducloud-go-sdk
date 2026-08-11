package blb

type CreateAppBlbServerGroupRequest struct {
	BlbId                   *string                      `json:"-"`
	ClientToken             *string                      `json:"-"`
	Name                    *string                      `json:"name,omitempty"`
	Desc                    *string                      `json:"desc,omitempty"`
	PreserveClientIpEnabled *bool                        `json:"preserveClientIpEnabled,omitempty"`
	BackendServerList       []*AppBackendServerForCreate `json:"backendServerList,omitempty"`
}
