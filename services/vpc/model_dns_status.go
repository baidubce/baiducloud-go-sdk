package vpc

type DnsStatus struct {
	Close   *string `json:"close,omitempty"`
	Wait    *string `json:"wait,omitempty"`
	Syncing *string `json:"syncing,omitempty"`
	Open    *string `json:"open,omitempty"`
	Closing *string `json:"closing,omitempty"`
}
