package ocr

type MultiIdcardItem struct {
	CardInfo   *MultiIdcardCardInfo `json:"card_info,omitempty"`
	CardResult *interface{}         `json:"card_result,omitempty"`
}
