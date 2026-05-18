package eip

type QueryTheListOfSharedTrafficPackagesRequest struct {
	Marker       *string `json:"-"`
	MaxKeys      *int32  `json:"-"`
	Id           *string `json:"-"`
	Status       *string `json:"-"`
	DeductPolicy *string `json:"-"`
}
