package privatezone

type DnsResolverRule struct {
	Id               *string            `json:"id,omitempty"`
	Name             *string            `json:"name,omitempty"`
	Status           *string            `json:"status,omitempty"`
	Description      *string            `json:"description,omitempty"`
	Zone             *string            `json:"zone,omitempty"`
	ResolverId       *string            `json:"resolverId,omitempty"`
	ResolverRegion   *string            `json:"resolverRegion,omitempty"`
	DnsServerConfigs []*DnsServerConfig `json:"dnsServerConfigs,omitempty"`
	CreateTime       *string            `json:"createTime,omitempty"`
	UpdateTime       *string            `json:"updateTime,omitempty"`
}
