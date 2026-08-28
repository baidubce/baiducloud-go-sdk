package ocr

type ImageResult struct {
	ImageId      *string           `json:"imageId,omitempty"`
	ImageUrl     *string           `json:"imageUrl,omitempty"`
	PaperSubject *string           `json:"paperSubject,omitempty"`
	ResizeRatio  *float64          `json:"resize_ratio,omitempty"`
	Result       []*QuestionResult `json:"result,omitempty"`
}
