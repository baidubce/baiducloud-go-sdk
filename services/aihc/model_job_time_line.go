package aihc

type JobTimeLine struct {
	ConditionType    *string `json:"conditionType,omitempty"`
	ConditionMessage *string `json:"conditionMessage,omitempty"`
	Time             *string `json:"time,omitempty"`
}
