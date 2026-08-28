package ocr

type HandwritingGetCharInfo struct {
	IsPunctuation *string             `json:"isPunctuation,omitempty"`
	Bbox          *HandwritingGetBBox `json:"bbox,omitempty"`
	Char          *string             `json:"char,omitempty"`
	Index         *string             `json:"index,omitempty"`
}
