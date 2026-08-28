package ocr

type PaperCutEduRequest struct {
	Image           *string `json:"image,omitempty"`
	Url             *string `json:"url,omitempty"`
	PdfFile         *string `json:"pdf_file,omitempty"`
	PdfFileNum      *int32  `json:"pdf_file_num,omitempty"`
	LanguageType    *string `json:"language_type,omitempty"`
	DetectDirection *bool   `json:"detect_direction,omitempty"`
	WordsType       *string `json:"words_type,omitempty"`
	SpliceText      *bool   `json:"splice_text,omitempty"`
	Enhance         *bool   `json:"enhance,omitempty"`
	OnlySplit       *bool   `json:"only_split,omitempty"`
}
