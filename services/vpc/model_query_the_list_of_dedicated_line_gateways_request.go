package vpc

type QueryTheListOfDedicatedLineGatewaysRequest struct {
	VpcId       *string `json:"-"`
	EtGatewayId *string `json:"-"`
	Name        *string `json:"-"`
	Status      *string `json:"-"`
	Marker      *string `json:"-"`
	MaxKeys     *int32  `json:"-"`
}
