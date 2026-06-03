package ccr

type Privatelinks struct {
	VpcID        *string `json:"vpcID,omitempty"`
	SubnetID     *string `json:"subnetID,omitempty"`
	ServiceNetID *string `json:"serviceNetID,omitempty"`
	IpAddress    *string `json:"ipAddress,omitempty"`
	Status       *string `json:"status,omitempty"`
}
