package ocr

type WeightNoteWordsResult struct {
	PlateNum         []*WeightNoteWordItem `json:"PlateNum,omitempty"`
	PrintTime        []*WeightNoteWordItem `json:"PrintTime,omitempty"`
	CrossWeight      []*WeightNoteWordItem `json:"CrossWeight,omitempty"`
	TareWeight       []*WeightNoteWordItem `json:"TareWeight,omitempty"`
	NetWeight        []*WeightNoteWordItem `json:"NetWeight,omitempty"`
	SendingCompany   []*WeightNoteWordItem `json:"SendingCompany,omitempty"`
	ReceivingCompany []*WeightNoteWordItem `json:"ReceivingCompany,omitempty"`
	DeliveryNumber   []*WeightNoteWordItem `json:"DeliveryNumber,omitempty"`
}
