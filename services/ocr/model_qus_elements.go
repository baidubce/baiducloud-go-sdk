package ocr

type QusElements struct {
	QuestionWords *string   `json:"question_words,omitempty"`
	Choices       []*string `json:"choices,omitempty"`
	QusType       *string   `json:"qus_type,omitempty"`
	AnswerWords   []*string `json:"answer_words,omitempty"`
}
