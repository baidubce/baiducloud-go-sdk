package ocr

type AccurateBasicWordsResult struct {
	Words       *string                   `json:"words,omitempty"`
	Probability *AccurateBasicProbability `json:"probability,omitempty"`
}
