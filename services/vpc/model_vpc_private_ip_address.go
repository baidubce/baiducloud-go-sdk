package vpc

type VpcPrivateIpAddress struct {
	Cidr                 *string `json:"cidr,omitempty"`
	PrivateIpAddress     *string `json:"privateIpAddress,omitempty"`
	PrivateIpAddressType *string `json:"privateIpAddressType,omitempty"`
	CreatedTime          *string `json:"createdTime,omitempty"`
}
