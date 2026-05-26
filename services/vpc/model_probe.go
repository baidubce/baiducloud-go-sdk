package vpc

type Probe struct {
	ProbeId     *string   `json:"probeId,omitempty"`
	Description *string   `json:"description,omitempty"`
	DestIp      *string   `json:"destIp,omitempty"`
	DestPort    *string   `json:"destPort,omitempty"`
	Frequency   *int32    `json:"frequency,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Payload     *string   `json:"payload,omitempty"`
	Protocol    *string   `json:"protocol,omitempty"`
	SourceIps   []*string `json:"sourceIps,omitempty"`
	Status      *string   `json:"status,omitempty"`
	SubnetId    *string   `json:"subnetId,omitempty"`
	VpcId       *string   `json:"vpcId,omitempty"`
}
