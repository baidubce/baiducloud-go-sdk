package ocr

type PaperCutEduVlmGetQusResult struct {
	QusId       *int32                  `json:"qus_id,omitempty"`
	Location    *PaperCutEduVlmLocation `json:"location,omitempty"`
	QusElements *QusElements            `json:"qus_elements,omitempty"`
}
