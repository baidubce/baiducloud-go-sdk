package face

type VerifyRequest struct {
	VideoBase64  *string `json:"video_base64,omitempty"`
	TypeIdentify *string `json:"type_identify,omitempty"`
	SessionId    *string `json:"session_id,omitempty"`
	LipIdentify  *string `json:"lip_identify,omitempty"`
	FaceField    *string `json:"face_field,omitempty"`
}
