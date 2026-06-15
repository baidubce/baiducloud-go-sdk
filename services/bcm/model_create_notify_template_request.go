package bcm

type CreateNotifyTemplateRequest struct {
	Name           *string           `json:"name,omitempty"`
	SilencePeriods []*SilencePeriod  `json:"silencePeriods,omitempty"`
	Receivers      []*NotifyReceiver `json:"receivers,omitempty"`
	Callbacks      []*Callback       `json:"callbacks,omitempty"`
}
