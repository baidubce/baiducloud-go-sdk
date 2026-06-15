package bcc

type AutoReleaseInstanceRequest struct {
	InstanceId             *string `json:"-"`
	IsEipAutoRelatedDelete *bool   `json:"isEipAutoRelatedDelete,omitempty"`
	ReleaseTime            *string `json:"releaseTime,omitempty"`
}
