package ocr

type QuestionResult struct {
	CorrectResult *int32  `json:"correctResult,omitempty"`
	QuestionId    *string `json:"questionId,omitempty"`
	Question      *string `json:"question,omitempty"`
	QuestionArea  []*Area `json:"questionArea,omitempty"`
	IsFinish      *bool   `json:"isFinish,omitempty"`
	Seqence       *int32  `json:"seqence,omitempty"`
	OcrType       *int32  `json:"type,omitempty"`
	CropUrl       *string `json:"cropUrl,omitempty"`
	Slot          []*Slot `json:"slot,omitempty"`
}
