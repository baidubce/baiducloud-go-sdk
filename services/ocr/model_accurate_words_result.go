package ocr

type AccurateWordsResult struct {
	Words                          *string                     `json:"words,omitempty"`
	Location                       *AccurateLocation           `json:"location,omitempty"`
	Chars                          []*AccurateChar             `json:"chars,omitempty"`
	Probability                    *AccurateProbability        `json:"probability,omitempty"`
	VertexesLocation               []*AccurateVertexesLocation `json:"vertexes_location,omitempty"`
	FinegrainedVertexesLocation    []*AccurateVertexesLocation `json:"finegrained_vertexes_location,omitempty"`
	MinFinegrainedVertexesLocation []*AccurateVertexesLocation `json:"min_finegrained_vertexes_location,omitempty"`
}
