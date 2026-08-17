package ocr

type MedicalRecordLocation struct {
	Top    *int32 `json:"top,omitempty"`
	Left   *int32 `json:"left,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Width  *int32 `json:"width,omitempty"`
}
