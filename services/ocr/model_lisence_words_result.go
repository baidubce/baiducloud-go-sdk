package ocr

type LisenceWordsResult struct {
	Color            *string           `json:"color,omitempty"`
	Number           *string           `json:"number,omitempty"`
	Probability      []*float32        `json:"probability,omitempty"`
	VertexesLocation []*VertexLocation `json:"vertexes_location,omitempty"`
	CoverInfo        *string           `json:"cover_info,omitempty"`
	EditTool         *string           `json:"edit_tool,omitempty"`
}
