package bls

type SetSingleLogShipperStatusRequest struct {
	LogShipperID  *string `json:"-"`
	DesiredStatus *string `json:"desiredStatus,omitempty"`
}
