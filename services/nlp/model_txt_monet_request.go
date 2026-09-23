package nlp

type TxtMonetRequest struct {
	Charset     *string        `json:"-"`
	ContentList []*ContentItem `json:"content_list,omitempty"`
}
