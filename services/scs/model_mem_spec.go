package scs

type MemSpec struct {
	MemUsageUpperThreshold        *int32  `json:"memUsageUpperThreshold,omitempty"`
	MemUsageDownThreshold         *int32  `json:"memUsageDownThreshold,omitempty"`
	MaxNodeType                   *string `json:"maxNodeType,omitempty"`
	MinNodeType                   *string `json:"minNodeType,omitempty"`
	ObservationWindowSizeForUpper *string `json:"observationWindowSizeForUpper,omitempty"`
	ObservationWindowSizeForDown  *string `json:"observationWindowSizeForDown,omitempty"`
}
