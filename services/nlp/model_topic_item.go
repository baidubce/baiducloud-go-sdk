package nlp

type TopicItem struct {
	Lv1TagList []*TopicTagItem `json:"lv1_tag_list,omitempty"`
	Lv2TagList []*TopicTagItem `json:"lv2_tag_list,omitempty"`
}
