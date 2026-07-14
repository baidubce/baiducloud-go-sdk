package as

type ZoneInfo struct {
	Zone       *string `json:"zone,omitempty"`
	SubnetId   *string `json:"subnetId,omitempty"`
	SubnetUuid *string `json:"subnetUuid,omitempty"`
	SubnetName *string `json:"subnetName,omitempty"`
	SubnetType *int32  `json:"subnetType,omitempty"`
	NodeCount  *int32  `json:"nodeCount,omitempty"`
}
