package bcc

type DelIpv6Request struct {
	InstanceId  *string `json:"instanceId,omitempty"`
	Ipv6Address *string `json:"ipv6Address,omitempty"`
	Reboot      *bool   `json:"reboot,omitempty"`
}
