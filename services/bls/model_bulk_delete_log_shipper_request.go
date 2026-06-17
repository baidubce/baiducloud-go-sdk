package bls

type BulkDeleteLogShipperRequest struct {
	LogShipperIDs []*string `json:"logShipperIDs,omitempty"`
}
