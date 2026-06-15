package bcc

type BatchAddIpRequest struct {
	InstanceId                     *string   `json:"instanceId,omitempty"`
	SecondaryPrivateIpAddressCount *int32    `json:"secondaryPrivateIpAddressCount,omitempty"`
	PrivateIps                     []*string `json:"privateIps,omitempty"`
	AllocateMultiIpv6Addr          *bool     `json:"allocateMultiIpv6Addr,omitempty"`
}
