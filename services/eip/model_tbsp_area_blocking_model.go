package eip

type TbspAreaBlockingModel struct {
	Ip             *string `json:"ip,omitempty"`
	BlockArea      *string `json:"blockArea,omitempty"`
	BlockBeginTime *string `json:"blockBeginTime,omitempty"`
	BlockEndTime   *string `json:"blockEndTime,omitempty"`
	BlockType      *string `json:"blockType,omitempty"`
}
