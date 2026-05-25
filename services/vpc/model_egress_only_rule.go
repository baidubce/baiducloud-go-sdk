package vpc

type EgressOnlyRule struct {
	EgressOnlyRuleId *string `json:"egressOnlyRuleId,omitempty"`
	Cidr             *string `json:"cidr,omitempty"`
}
