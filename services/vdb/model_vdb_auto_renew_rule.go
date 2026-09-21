package vdb

type VdbAutoRenewRule struct {
	RenewTime     *int32  `json:"renewTime,omitempty"`
	RenewTimeUnit *string `json:"renewTimeUnit,omitempty"`
}
