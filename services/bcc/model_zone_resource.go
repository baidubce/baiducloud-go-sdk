package bcc

type ZoneResource struct {
	ZoneName     *string            `json:"zoneName,omitempty"`
	BccResources []*BccBidResources `json:"bccResources,omitempty"`
}
