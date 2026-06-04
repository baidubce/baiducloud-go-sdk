package privatezone

type UpdateResolverRuleRequest struct {
	ReluId           *string            `json:"-"`
	ClientToken      *string            `json:"-"`
	Name             *string            `json:"name,omitempty"`
	Description      *string            `json:"description,omitempty"`
	DnsServerConfigs []*DnsServerConfig `json:"dnsServerConfigs,omitempty"`
}
