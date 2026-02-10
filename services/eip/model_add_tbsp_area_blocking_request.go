package eip

type AddTbspAreaBlockingRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	Ip          *string `json:"ip,omitempty"`
	BlockTime   *int32  `json:"blockTime,omitempty"`
	BlockType   *string `json:"blockType,omitempty"`
}
