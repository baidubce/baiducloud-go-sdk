package ocr

type Word struct {
	Word          *string        `json:"word,omitempty"`
	WordsLocation *WordsLocation `json:"words_location,omitempty"`
}
