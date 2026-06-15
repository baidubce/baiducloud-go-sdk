package bcc

type PrepayConfig struct {
	InstanceId      *string   `json:"instanceId,omitempty"`
	AutoRenew       *bool     `json:"autoRenew,omitempty"`
	AutoRenewPeriod *int32    `json:"autoRenewPeriod,omitempty"`
	Duration        *int32    `json:"duration,omitempty"`
	CdsList         []*string `json:"cdsList,omitempty"`
	AutoPay         *bool     `json:"autoPay,omitempty"`
}
