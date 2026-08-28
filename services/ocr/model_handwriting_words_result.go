package ocr

type HandwritingWordsResult struct {
	Location    *HandWritingLocation    `json:"location,omitempty"`
	Words       *string                 `json:"words,omitempty"`
	Chars       []*HandwritingChar      `json:"chars,omitempty"`
	Probability *HandWritingProbability `json:"probability,omitempty"`
}
