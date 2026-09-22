package face

type Gender struct {
	FaceType    *string  `json:"type,omitempty"`
	Probability *float64 `json:"probability,omitempty"`
}
