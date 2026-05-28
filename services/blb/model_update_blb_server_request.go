package blb

type UpdateBlbServerRequest struct {
	BlbId             *string               `json:"-"`
	ClientToken       *string               `json:"-"`
	BackendServerList []*BackendServerModel `json:"backendServerList,omitempty"`
}
