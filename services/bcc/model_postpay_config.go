package bcc

type PostpayConfig struct {
	InstanceId    *string   `json:"instanceId,omitempty"`
	CdsList       []*string `json:"cdsList,omitempty"`
	EffectiveType *string   `json:"effectiveType,omitempty"`
}
