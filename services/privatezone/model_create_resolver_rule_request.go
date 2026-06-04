package privatezone

type CreateResolverRuleRequest struct {
	ClientToken      *string            `json:"-"`
	Name             *string            `json:"name,omitempty"`
	Description      *string            `json:"description,omitempty"`
	Zone             *string            `json:"zone,omitempty"`
	ResolverId       *string            `json:"resolverId,omitempty"`
	DnsServerConfigs []*DnsServerConfig `json:"dnsServerConfigs,omitempty"`
}
