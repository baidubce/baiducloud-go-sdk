package bls

type ConditionValidateResult struct {
	Valid   *bool   `json:"valid,omitempty"`
	Message *string `json:"message,omitempty"`
}
