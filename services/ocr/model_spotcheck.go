package ocr

type Spotcheck struct {
	No           *string `json:"no,omitempty"`
	Executiveorg *string `json:"executiveorg,omitempty"`
	OcrType      *string `json:"type,omitempty"`
	Date         *string `json:"date,omitempty"`
	Consequence  *string `json:"consequence,omitempty"`
	Remark       *string `json:"remark,omitempty"`
}
