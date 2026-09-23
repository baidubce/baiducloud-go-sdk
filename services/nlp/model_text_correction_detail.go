package nlp

type TextCorrectionDetail struct {
	SentenceId           *int32                       `json:"sentence_id,omitempty"`
	Sentence             *string                      `json:"sentence,omitempty"`
	SentenceFixed        *string                      `json:"sentence_fixed,omitempty"`
	VecFragment          []*TextCorrectionVecFragment `json:"vec_fragment,omitempty"`
	BeginSentenceOffset  *int32                       `json:"begin_sentence_offset,omitempty"`
	EndSentenceOffset    *int32                       `json:"end_sentence_offset,omitempty"`
	BeginPsentContOffset *int32                       `json:"begin_psent_cont_offset,omitempty"`
	EndPsentContOffset   *int32                       `json:"end_psent_cont_offset,omitempty"`
}
