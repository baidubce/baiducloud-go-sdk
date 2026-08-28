package ocr

type WeightNoteWordItem struct {
	Word        *string      `json:"word,omitempty"`
	Probability *interface{} `json:"probability,omitempty"`
}
