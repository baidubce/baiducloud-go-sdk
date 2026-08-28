package ocr

type GeneralParagraphsResult struct {
	WordsResultIdx                 []*int32                   `json:"words_result_idx,omitempty"`
	FinegrainedVertexesLocation    []*GeneralVertexesLocation `json:"finegrained_vertexes_location,omitempty"`
	MinFinegrainedVertexesLocation []*GeneralVertexesLocation `json:"min_finegrained_vertexes_location,omitempty"`
}
