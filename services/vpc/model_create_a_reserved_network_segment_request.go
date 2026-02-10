package vpc

type CreateAReservedNetworkSegmentRequest struct {
	ClientToken *string `json:"-"`
	SubnetId    *string `json:"subnetId,omitempty"`
	IpCidr      *string `json:"ipCidr,omitempty"`
	IpVersion   *int32  `json:"ipVersion,omitempty"`
	Description *string `json:"description,omitempty"`
}
