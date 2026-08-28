package ocr

type WebImageLocWordsResult struct {
	Words        *string                    `json:"words,omitempty"`
	Location     *WebImageLocLocation       `json:"location,omitempty"`
	Probability  *WebImageLocProbability    `json:"probability,omitempty"`
	PolyLocation []*WebImageLocPolyLocation `json:"poly_location,omitempty"`
	Chars        []*WebImageLocChar         `json:"chars,omitempty"`
}
