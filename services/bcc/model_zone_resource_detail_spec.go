package bcc

type ZoneResourceDetailSpec struct {
	ZoneName     *string       `json:"zoneName,omitempty"`
	EbcResources *BbcResources `json:"ebcResources,omitempty"`
	BccResources *BccResources `json:"bccResources,omitempty"`
}
