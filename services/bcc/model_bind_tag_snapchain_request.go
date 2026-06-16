package bcc

type BindTagSnapchainRequest struct {
	ChainId    *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
