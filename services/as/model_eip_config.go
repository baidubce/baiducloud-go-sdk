package as

type EipConfig struct {
	EipGroupBindStrategy   *string            `json:"eipGroupBindStrategy,omitempty"`
	EipGroupUnbindStrategy *string            `json:"eipGroupUnbindStrategy,omitempty"`
	EipGroupIdList         []*string          `json:"eipGroupIdList,omitempty"`
	Increase               *EipGroupIncrease  `json:"increase,omitempty"`
	Decrease               *EipGroupDecrease  `json:"decrease,omitempty"`
	Bandwidth              *EipGroupBandwidth `json:"bandwidth,omitempty"`
}
