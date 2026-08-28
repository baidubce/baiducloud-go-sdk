package ocr

type AccurateChar struct {
	Char     *string           `json:"char,omitempty"`
	CharProb *int32            `json:"char_prob,omitempty"`
	Location *AccurateLocation `json:"location,omitempty"`
}
