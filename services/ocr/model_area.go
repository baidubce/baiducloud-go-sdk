package ocr

type Area struct {
	LeftX  *int32 `json:"left_x,omitempty"`
	LeftY  *int32 `json:"left_y,omitempty"`
	RightX *int32 `json:"right_x,omitempty"`
	RightY *int32 `json:"right_y,omitempty"`
}
