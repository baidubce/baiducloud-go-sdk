package eip

type QueryTheListOfEipsInTheRecyclingBinRequest struct {
	Eip     *string `json:"-"`
	Name    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
