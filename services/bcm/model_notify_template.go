package bcm

type NotifyTemplate struct {
	Id             *string           `json:"id,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Source         *string           `json:"source,omitempty"`
	SilencePeriods []*SilencePeriod  `json:"silencePeriods,omitempty"`
	Receivers      []*NotifyReceiver `json:"receivers,omitempty"`
	Callbacks      []*Callback       `json:"callbacks,omitempty"`
}
