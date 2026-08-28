package ocr

type MeterWordsResult struct {
	Words        *string           `json:"words,omitempty"`
	Location     *MeterLocation    `json:"location,omitempty"`
	Probability  *MeterProbability `json:"probability,omitempty"`
	PolyLocation []*MeterPolyPoint `json:"poly_location,omitempty"`
}
