package as

type ExpectAction struct {
	ActionType  *string `json:"actionType,omitempty"`
	ActionNum   *int32  `json:"actionNum,omitempty"`
	AdjustToNum *int32  `json:"adjustToNum,omitempty"`
}
