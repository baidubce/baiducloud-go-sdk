package rds

type InstanceEnablesAutomaticRenewalRequest struct {
	InstanceIds       []*string `json:"instanceIds,omitempty"`
	AutoRenewTimeUnit *string   `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     *int32    `json:"autoRenewTime,omitempty"`
}
