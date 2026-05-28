package blb

type AddBlbServerRequest struct {
	BlbId             *string               `json:"-"`
	ClientToken       *string               `json:"-"`
	BackendServerList []*BackendServerModel `json:"backendServerList,omitempty"`
}
