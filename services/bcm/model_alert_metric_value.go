package bcm

type AlertMetricValue struct {
	Name  *string  `json:"name,omitempty"`
	Label *string  `json:"label,omitempty"`
	Value *float32 `json:"value,omitempty"`
	Unit  *string  `json:"unit,omitempty"`
}
