package ocr

type PaperCutEduVlmLocation struct {
	QusLocation []*float64 `json:"qus_location,omitempty"`
	PicLocation []*float64 `json:"pic_location,omitempty"`
	AnsLocation []*float64 `json:"ans_location,omitempty"`
}
