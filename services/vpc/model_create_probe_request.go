package vpc

type CreateProbeRequest struct {
	ClientToken *string   `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	VpcId       *string   `json:"vpcId,omitempty"`
	SubnetId    *string   `json:"subnetId,omitempty"`
	Protocol    *string   `json:"protocol,omitempty"`
	Frequency   *int32    `json:"frequency,omitempty"`
	SourceIps   []*string `json:"sourceIps,omitempty"`
	SourceIpNum *int32    `json:"sourceIpNum,omitempty"`
	DestIp      *string   `json:"destIp,omitempty"`
	DestPort    *int32    `json:"destPort,omitempty"`
	Payload     *string   `json:"payload,omitempty"`
}
