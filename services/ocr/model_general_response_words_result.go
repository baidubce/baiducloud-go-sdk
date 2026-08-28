package ocr

type GeneralResponseWordsResult struct {
	Words                          *string                    `json:"words,omitempty"`
	Location                       *GeneralLocation           `json:"location,omitempty"`
	Chars                          []*GeneralResponseChar     `json:"chars,omitempty"`
	Probability                    *GeneralProbability        `json:"probability,omitempty"`
	VertexesLocation               []*GeneralVertexesLocation `json:"vertexes_location,omitempty"`
	FinegrainedVertexesLocation    []*GeneralVertexesLocation `json:"finegrained_vertexes_location,omitempty"`
	MinFinegrainedVertexesLocation []*GeneralVertexesLocation `json:"min_finegrained_vertexes_location,omitempty"`
}
