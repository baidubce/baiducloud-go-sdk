package face

type MaskInfo struct {
	FaceType    *int32   `json:"type,omitempty"`
	Probability *float64 `json:"probability,omitempty"`
}
