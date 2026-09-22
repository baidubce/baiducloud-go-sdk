package face

type FacePersonVerifyRequest struct {
	Image           *string `json:"image,omitempty"`
	ImageType       *string `json:"image_type,omitempty"`
	IdCardNumber    *string `json:"id_card_number,omitempty"`
	Name            *string `json:"name,omitempty"`
	QualityControl  *string `json:"quality_control,omitempty"`
	LivenessControl *string `json:"liveness_control,omitempty"`
	SpoofingControl *string `json:"spoofing_control,omitempty"`
}
