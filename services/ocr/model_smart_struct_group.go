package ocr

type SmartStructGroup struct {
	Key   []*SmartStructTextLine `json:"key,omitempty"`
	Value []*SmartStructTextLine `json:"value,omitempty"`
}
