package eip

type UpdateEipBpAutoReleaseTimeRequest struct {
	Id              *string `json:"-"`
	ClientToken     *string `json:"-"`
	AutoReleaseTime *string `json:"autoReleaseTime,omitempty"`
}
