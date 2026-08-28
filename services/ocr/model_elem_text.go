package ocr

type ElemText struct {
	StemText           *string `json:"stem_text,omitempty"`
	SubqusText         *string `json:"subqus_text,omitempty"`
	AnswerText         *string `json:"answer_text,omitempty"`
	OptionText         *string `json:"option_text,omitempty"`
	InterpretationText *string `json:"interpretation_text,omitempty"`
}
