package ocr

type Underline struct {
	Points *Points  `json:"points,omitempty"`
	Prob   *float64 `json:"prob,omitempty"`
}
