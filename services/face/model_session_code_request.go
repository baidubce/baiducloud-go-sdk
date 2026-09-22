package face

type SessionCodeRequest struct {
	FaceType      *string `json:"type,omitempty"`
	MinCodeLength *string `json:"min_code_length,omitempty"`
	MaxCodeLength *string `json:"max_code_length,omitempty"`
}
