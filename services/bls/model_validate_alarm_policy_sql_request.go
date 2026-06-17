package bls

type ValidateAlarmPolicySQLRequest struct {
	LogStores []*LogStore `json:"logStores,omitempty"`
	Query     *string     `json:"query,omitempty"`
}
