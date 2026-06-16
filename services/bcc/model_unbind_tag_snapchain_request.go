package bcc

type UnbindTagSnapchainRequest struct {
	ChainId    *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
