package blb

type DeleteBlbServerRequest struct {
	BlbId             *string   `json:"-"`
	ClientToken       *string   `json:"-"`
	BackendServerList []*string `json:"backendServerList,omitempty"`
}
