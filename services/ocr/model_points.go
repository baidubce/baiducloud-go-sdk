package ocr

type Points struct {
	StartX *int32 `json:"start_x,omitempty"`
	StartY *int32 `json:"start_y,omitempty"`
	EndX   *int32 `json:"end_x,omitempty"`
	EndY   *int32 `json:"end_y,omitempty"`
}
