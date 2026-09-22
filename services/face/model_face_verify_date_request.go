package face

type FaceVerifyDateRequest struct {
	Name            *string `json:"name,omitempty"`
	IdCardNumber    *string `json:"id_card_number,omitempty"`
	StartDate       *string `json:"start_date,omitempty"`
	EndDate         *string `json:"end_date,omitempty"`
	Image           *string `json:"image,omitempty"`
	ImageType       *string `json:"image_type,omitempty"`
	LivenessControl *string `json:"liveness_control,omitempty"`
	SpoofingControl *string `json:"spoofing_control,omitempty"`
	QualityControl  *string `json:"quality_control,omitempty"`
}
