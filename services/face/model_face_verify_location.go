package face

type FaceVerifyLocation struct {
	Left     *float64 `json:"left,omitempty"`
	Top      *float64 `json:"top,omitempty"`
	Width    *float64 `json:"width,omitempty"`
	Height   *float64 `json:"height,omitempty"`
	Rotation *int64   `json:"rotation,omitempty"`
}
