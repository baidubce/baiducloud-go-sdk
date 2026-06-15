package bcc

type ResourceBo struct {
	Id            *string     `json:"id,omitempty"`
	SerialNumber  *string     `json:"serialNumber,omitempty"`
	Name          *string     `json:"name,omitempty"`
	RecycleTime   *string     `json:"recycleTime,omitempty"`
	DeleteTime    *string     `json:"deleteTime,omitempty"`
	PaymentTiming *string     `json:"paymentTiming,omitempty"`
	ServiceName   *string     `json:"serviceName,omitempty"`
	ServiceType   *string     `json:"serviceType,omitempty"`
	ConfigItem    *ConfigItem `json:"configItem,omitempty"`
	ConfigItems   []*string   `json:"configItems,omitempty"`
}
