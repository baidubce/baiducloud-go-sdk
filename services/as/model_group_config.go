package as

type GroupConfig struct {
	MinNodeNum    *int32 `json:"minNodeNum,omitempty"`
	MaxNodeNum    *int32 `json:"maxNodeNum,omitempty"`
	CooldownInSec *int32 `json:"cooldownInSec,omitempty"`
	ExpectNum     *int32 `json:"expectNum,omitempty"`
	InitNum       *int32 `json:"initNum,omitempty"`
}
