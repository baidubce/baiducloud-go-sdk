package face

type FaceDetectTypeProbability struct {
	FaceType    *string  `json:"type,omitempty"`
	Probability *float64 `json:"probability,omitempty"`
}
