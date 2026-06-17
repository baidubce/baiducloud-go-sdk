package bls

type ValidateAlarmConditionRequest struct {
	FieldTypes []*string `json:"fieldTypes,omitempty"`
	Conditions []*string `json:"conditions,omitempty"`
}
