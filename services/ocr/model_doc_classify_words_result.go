package ocr

type DocClassifyWordsResult struct {
	OcrType    *string              `json:"type,omitempty"`
	Probablity *float32             `json:"probablity,omitempty"`
	Location   *DocClassifyLocation `json:"location,omitempty"`
}
