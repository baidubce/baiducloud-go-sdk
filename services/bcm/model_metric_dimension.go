package bcm

type MetricDimension struct {
	Key      *string   `json:"key,omitempty"`
	Operator *string   `json:"operator,omitempty"`
	Values   []*string `json:"values,omitempty"`
}
