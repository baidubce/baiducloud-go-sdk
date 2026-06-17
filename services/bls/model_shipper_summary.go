package bls

type ShipperSummary struct {
	LogShipperID   *string `json:"logShipperID,omitempty"`
	LogShipperName *string `json:"logShipperName,omitempty"`
	Project        *string `json:"project,omitempty"`
	LogStoreName   *string `json:"logStoreName,omitempty"`
	DestType       *string `json:"destType,omitempty"`
	Status         *string `json:"status,omitempty"`
	CreateDateTime *string `json:"createDateTime,omitempty"`
	ErrMessage     *string `json:"errMessage,omitempty"`
}
