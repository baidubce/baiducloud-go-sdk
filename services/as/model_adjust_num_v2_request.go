package as

type AdjustNumV2Request struct {
	GroupId    *string `json:"-"`
	AdjustNode *string `json:"-"`
	AdjustNum  *int32  `json:"adjustNum,omitempty"`
}
