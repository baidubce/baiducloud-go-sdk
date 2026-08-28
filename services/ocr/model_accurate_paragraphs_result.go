package ocr

type AccurateParagraphsResult struct {
	WordsResultIdx                 []*int32                    `json:"words_result_idx,omitempty"`
	FinegrainedVertexesLocation    []*AccurateVertexesLocation `json:"finegrained_vertexes_location,omitempty"`
	MinFinegrainedVertexesLocation []*AccurateVertexesLocation `json:"min_finegrained_vertexes_location,omitempty"`
}
