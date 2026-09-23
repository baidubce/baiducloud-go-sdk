package nlp

type ContentItem struct {
	Content   *string      `json:"content,omitempty"`
	QueryList []*QueryItem `json:"query_list,omitempty"`
}
