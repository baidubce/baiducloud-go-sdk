package nlp

type TextCorrectionVecFragment struct {
	OriFrag          *string  `json:"ori_frag,omitempty"`
	CorrectFrag      *string  `json:"correct_frag,omitempty"`
	BeginPos         *int32   `json:"begin_pos,omitempty"`
	EndPos           *int32   `json:"end_pos,omitempty"`
	Explain          *string  `json:"explain,omitempty"`
	ExplainLong      *string  `json:"explain_long,omitempty"`
	ExplainStructure *string  `json:"explain_structure,omitempty"`
	Operation        *int32   `json:"operation,omitempty"`
	Label            *string  `json:"label,omitempty"`
	Score            *float64 `json:"score,omitempty"`
}
