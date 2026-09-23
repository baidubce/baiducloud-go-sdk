package nlp

type VecFragment struct {
	OriFrag     *string `json:"ori_frag,omitempty"`
	CorrectFrag *string `json:"correct_frag,omitempty"`
	BeginPos    *int32  `json:"begin_pos,omitempty"`
	EndPos      *int32  `json:"end_pos,omitempty"`
}
