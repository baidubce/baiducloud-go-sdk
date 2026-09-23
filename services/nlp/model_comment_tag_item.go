package nlp

type CommentTagItem struct {
	Prop      *string `json:"prop,omitempty"`
	Adj       *string `json:"adj,omitempty"`
	Sentiment *int32  `json:"sentiment,omitempty"`
	BeginPos  *int32  `json:"begin_pos,omitempty"`
	EndPos    *int32  `json:"end_pos,omitempty"`
	Abstract  *string `json:"abstract,omitempty"`
}
