package as

type AsAlarmRule struct {
	Id                  *int32         `json:"id,omitempty"`
	Scope               *string        `json:"scope,omitempty"`
	MonitorObject       *MonitorObject `json:"monitorObject,omitempty"`
	Rules               [][]*AlarmRule `json:"rules,omitempty"`
	AlarmName           *string        `json:"alarmName,omitempty"`
	AliasName           *string        `json:"aliasName,omitempty"`
	InsufficientCycle   *int32         `json:"insufficientCycle,omitempty"`
	PolicyEnabled       *bool          `json:"policyEnabled,omitempty"`
	RuleContents        []*string      `json:"ruleContents,omitempty"`
	RuleContentsEn      []*string      `json:"ruleContentsEn,omitempty"`
	Source              *string        `json:"source,omitempty"`
	ComponentType       *string        `json:"componentType,omitempty"`
	AlarmActions        []*string      `json:"alarmActions,omitempty"`
	OkActions           []*string      `json:"okActions,omitempty"`
	InsufficientActions []*string      `json:"insufficientActions,omitempty"`
}
