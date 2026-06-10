package bcc

type ModifyVolumeChargeTypeRequest struct {
	VolumeId      *string  `json:"-"`
	Billing       *Billing `json:"billing,omitempty"`
	EffectiveType *string  `json:"effectiveType,omitempty"`
}
