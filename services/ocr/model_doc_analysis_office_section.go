package ocr

type DocAnalysisOfficeSection struct {
	Attribute     *string                  `json:"attribute,omitempty"`
	SectionsProb  *float64                 `json:"sections_prob,omitempty"`
	AttriLocation *AttriLocation           `json:"attri_location,omitempty"`
	SecIdx        *DocAnalysisOfficeSecIdx `json:"sec_idx,omitempty"`
}
