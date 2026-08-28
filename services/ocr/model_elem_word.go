package ocr

type ElemWord struct {
	WordLocation *WordLocation `json:"word_location,omitempty"`
	WordType     *string       `json:"word_type,omitempty"`
	Word         *string       `json:"word,omitempty"`
}
