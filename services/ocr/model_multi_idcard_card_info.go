package ocr

type MultiIdcardCardInfo struct {
	CardLocation     *MultiIdcardLocation    `json:"card_location,omitempty"`
	CardType         *string                 `json:"card_type,omitempty"`
	Direction        *int32                  `json:"direction,omitempty"`
	ImageStatus      *string                 `json:"image_status,omitempty"`
	RiskType         *string                 `json:"risk_type,omitempty"`
	EditTool         *string                 `json:"edit_tool,omitempty"`
	CardQuality      *MultiIdcardCardQuality `json:"card_quality,omitempty"`
	Photo            *string                 `json:"photo,omitempty"`
	PhotoLocation    *MultiIdcardLocation    `json:"photo_location,omitempty"`
	CardImage        *string                 `json:"card_image,omitempty"`
	IdcardNumberType *int32                  `json:"idcard_number_type,omitempty"`
}
