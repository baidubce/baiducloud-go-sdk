package ocr

type Slot struct {
	SlotId          *string `json:"slotId,omitempty"`
	Seqence         *int32  `json:"seqence,omitempty"`
	HandwritingArea *Area   `json:"handwritingArea,omitempty"`
	CorrectResult   *int32  `json:"correctResult,omitempty"`
	Reason          *string `json:"reason,omitempty"`
}
