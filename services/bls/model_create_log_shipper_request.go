package bls

type CreateLogShipperRequest struct {
	Project        *string     `json:"project,omitempty"`
	LogStoreName   *string     `json:"logStoreName,omitempty"`
	LogShipperName *string     `json:"logShipperName,omitempty"`
	StartTime      *string     `json:"startTime,omitempty"`
	DestType       *string     `json:"destType,omitempty"`
	DestConfig     *DestConfig `json:"destConfig,omitempty"`
}
