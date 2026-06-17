package bls

type BulkSetLogShipperStatusRequest struct {
	DesiredStatus *string   `json:"desiredStatus,omitempty"`
	LogShipperIDs []*string `json:"logShipperIDs,omitempty"`
}
