package vpc

type QueryTheListOfElasticNetworkCardsRequest struct {
	VpcId            *string   `json:"-"`
	InstanceId       *string   `json:"-"`
	Name             *string   `json:"-"`
	PrivateIpAddress []*string `json:"-"`
	Marker           *string   `json:"-"`
	MaxKeys          *int32    `json:"-"`
}
