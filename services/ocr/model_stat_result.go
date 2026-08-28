package ocr

type StatResult struct {
	All        *int32 `json:"all,omitempty"`
	Corrected  *int32 `json:"corrected,omitempty"`
	Correcting *int32 `json:"correcting,omitempty"`
}
