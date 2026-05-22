package cfw

type ListProtectInstancesRequest struct {
	InstanceType *string `json:"-"`
	Marker       *string `json:"-"`
	MaxKeys      *int32  `json:"-"`
	Status       *string `json:"-"`
	Region       *string `json:"-"`
}
