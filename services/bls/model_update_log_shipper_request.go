package bls

type UpdateLogShipperRequest struct {
	LogShipperID   *string     `json:"-"`
	Project        *string     `json:"project,omitempty"`
	LogShipperName *string     `json:"logShipperName,omitempty"`
	DestConfig     *DestConfig `json:"destConfig,omitempty"`
}
