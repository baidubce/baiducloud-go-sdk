package ocr

type GeneralBasicWordsResult struct {
	Words       *string                  `json:"words,omitempty"`
	Probability *GeneralBasicProbability `json:"probability,omitempty"`
}
