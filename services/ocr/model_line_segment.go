package ocr

type LineSegment struct {
	LineId     *string `json:"lineId,omitempty"`
	StartIndex *int32  `json:"startIndex,omitempty"`
	EndIndex   *int32  `json:"endIndex,omitempty"`
}
