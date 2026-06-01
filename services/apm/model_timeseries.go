package apm

type Timeseries struct {
	Metric     *string       `json:"metric,omitempty"`
	Dimensions []*Dimension  `json:"dimensions,omitempty"`
	Data       [][]float32   `json:"data,omitempty"`
	CompareTo  [][][]float32 `json:"compareTo,omitempty"`
}
