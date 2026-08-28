package ocr

type TablesResult struct {
	TableLocation []*TablePoint `json:"table_location,omitempty"`
	Header        []*Header     `json:"header,omitempty"`
	Body          []*Body       `json:"body,omitempty"`
	Footer        []*Footer     `json:"footer,omitempty"`
}
