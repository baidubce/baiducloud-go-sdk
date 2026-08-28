package ocr

type PaperCutEduVlmCreateTaskRequest struct {
	Image      *string `json:"image,omitempty"`
	Url        *string `json:"url,omitempty"`
	PdfFile    *string `json:"pdf_file,omitempty"`
	PdfFileNum *int32  `json:"pdf_file_num,omitempty"`
	OnlySplit  *bool   `json:"only_split,omitempty"`
	SceneType  *string `json:"scene_type,omitempty"`
	Enhance    *bool   `json:"enhance,omitempty"`
}
