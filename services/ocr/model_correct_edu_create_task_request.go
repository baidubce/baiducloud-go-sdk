package ocr

type CorrectEduCreateTaskRequest struct {
	Image             *string `json:"image,omitempty"`
	Url               *string `json:"url,omitempty"`
	OnlySplit         *bool   `json:"only_split,omitempty"`
	DisablePreprocess *bool   `json:"disable_preprocess,omitempty"`
}
