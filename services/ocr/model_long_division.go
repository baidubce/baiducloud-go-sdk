package ocr

type LongDivision struct {
	Location *DocAnalysisLocation `json:"location,omitempty"`
	Words    []*Word              `json:"words,omitempty"`
}
