package ocr

type DocAnalysisWords struct {
	Word            *string                     `json:"word,omitempty"`
	WordsLocation   *WordsLocation              `json:"words_location,omitempty"`
	PolyLocation    *DocAnalysisPolyLocation    `json:"poly_location,omitempty"`
	LineProbability *DocAnalysisLineProbability `json:"line_probability,omitempty"`
}
