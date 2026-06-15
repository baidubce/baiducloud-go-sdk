package bcc

type CdsCustomPeriod struct {
	Period   *int32  `json:"period,omitempty"`
	VolumeId *string `json:"volumeId,omitempty"`
}
