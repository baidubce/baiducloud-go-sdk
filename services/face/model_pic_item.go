package face

type PicItem struct {
	FaceToken *string  `json:"face_token,omitempty"`
	Spoofing  *float32 `json:"spoofing,omitempty"`
}
