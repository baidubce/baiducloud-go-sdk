package ocr

type TableResult struct {
	TableLocation []*DocAnalysisOfficePoint `json:"table_location,omitempty"`
	Header        []*TableHeader            `json:"header,omitempty"`
	Body          []*TableBody              `json:"body,omitempty"`
	Footer        []*TableFooter            `json:"footer,omitempty"`
}
