package ocr

type DocAnalysisSection struct {
	Attribute     *string                  `json:"attribute,omitempty"`
	AttriLocation *DocAnalysisPolyLocation `json:"attri_location,omitempty"`
	SecIdx        *DocAnalysisSecIdx       `json:"sec_idx,omitempty"`
}
