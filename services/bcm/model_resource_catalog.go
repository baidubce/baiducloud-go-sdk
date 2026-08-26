package bcm

type ResourceCatalog struct {
	Scope      *string                `json:"scope,omitempty"`
	ScopeLabel *string                `json:"scopeLabel,omitempty"`
	Resources  []*ResourceCatalogItem `json:"resources,omitempty"`
	Regions    []*string              `json:"regions,omitempty"`
}
