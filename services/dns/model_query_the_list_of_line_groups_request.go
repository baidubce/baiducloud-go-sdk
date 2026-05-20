package dns

type QueryTheListOfLineGroupsRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
