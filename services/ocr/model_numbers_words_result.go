package ocr

type NumbersWordsResult struct {
	Location *NumbersLocation `json:"location,omitempty"`
	Words    *string          `json:"words,omitempty"`
	Chars    []*NumberChar    `json:"chars,omitempty"`
}
