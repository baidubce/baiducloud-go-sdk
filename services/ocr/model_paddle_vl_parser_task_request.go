package ocr

type PaddleVlParserTaskRequest struct {
	FileData        *string `json:"file_data,omitempty"`
	FileUrl         *string `json:"file_url,omitempty"`
	FileName        *string `json:"file_name,omitempty"`
	AnalysisChart   *bool   `json:"analysis_chart,omitempty"`
	MergeTables     *bool   `json:"merge_tables,omitempty"`
	RelevelTitles   *bool   `json:"relevel_titles,omitempty"`
	RecognizeSeal   *bool   `json:"recognize_seal,omitempty"`
	ReturnSpanBoxes *bool   `json:"return_span_boxes,omitempty"`
}
