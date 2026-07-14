package as

type AsRecord struct {
	GroupId              *string         `json:"groupId,omitempty"`
	RecordId             *string         `json:"recordId,omitempty"`
	StartTime            *string         `json:"startTime,omitempty"`
	Result               *string         `json:"result,omitempty"`
	ActualScaleNode      []*string       `json:"actualScaleNode,omitempty"`
	ActualScaleBandwidth *int32          `json:"actualScaleBandwidth,omitempty"`
	CurrentBandwidth     *int32          `json:"currentBandwidth,omitempty"`
	RemainedNode         []*string       `json:"remainedNode,omitempty"`
	Action               *string         `json:"action,omitempty"`
	ScaleCondition       *ScaleCondition `json:"scaleCondition,omitempty"`
	RuleId               *string         `json:"ruleId,omitempty"`
	Message              *string         `json:"message,omitempty"`
	ExpectAction         *ExpectAction   `json:"expectAction,omitempty"`
	ExecuteType          *string         `json:"executeType,omitempty"`
	DagId                *string         `json:"dagId,omitempty"`
	Resource             *Resource       `json:"resource,omitempty"`
}
