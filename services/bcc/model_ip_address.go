package bcc

type IpAddress struct {
	Primary          *bool   `json:"primary,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	Ipv6Address      *string `json:"ipv6Address,omitempty"`
}
