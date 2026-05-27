package csn

type CreateCsnBpRequest struct {
	ClientToken   *string     `json:"-"`
	Name          *string     `json:"name,omitempty"`
	InterworkType *string     `json:"interworkType,omitempty"`
	Bandwidth     *int32      `json:"bandwidth,omitempty"`
	GeographicA   *string     `json:"geographicA,omitempty"`
	GeographicB   *string     `json:"geographicB,omitempty"`
	Billing       *Billing    `json:"billing,omitempty"`
	Tags          []*TagModel `json:"tags,omitempty"`
}
