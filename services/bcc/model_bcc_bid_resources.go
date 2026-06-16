package bcc

type BccBidResources struct {
	InstanceType *string          `json:"instanceType,omitempty"`
	Flavors      []*BccBidFlavors `json:"flavors,omitempty"`
}
