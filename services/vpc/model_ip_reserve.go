package vpc

type IpReserve struct {
	IpReserveId *string `json:"ipReserveId,omitempty"`
	SubnetId    *string `json:"subnetId,omitempty"`
	IpCidr      *string `json:"ipCidr,omitempty"`
	IpVersion   *string `json:"ipVersion,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
	UpdatedTime *string `json:"updatedTime,omitempty"`
}
