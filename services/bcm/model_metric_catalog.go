package bcm

type MetricCatalog struct {
	Name     *string          `json:"name,omitempty"`
	Label    *string          `json:"label,omitempty"`
	Catalogs []*MetricCatalog `json:"catalogs,omitempty"`
	Metrics  []*Metric        `json:"metrics,omitempty"`
}
