package bls

type SetSingleLogShipperStatusRequest struct {
	LogShipperID *string `json:"-"`
	Status       *string `json:"status,omitempty"`
}
