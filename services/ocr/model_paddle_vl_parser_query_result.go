package ocr

type PaddleVlParserQueryResult struct {
	TaskId         *string `json:"task_id,omitempty"`
	Status         *string `json:"status,omitempty"`
	TaskError      *string `json:"task_error,omitempty"`
	MarkdownUrl    *string `json:"markdown_url,omitempty"`
	ParseResultUrl *string `json:"parse_result_url,omitempty"`
}
