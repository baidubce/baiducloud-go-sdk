package image

type VehicleInfo struct {
	Location    *VehicleDetectLocation `json:"location,omitempty"`
	ImageType   *string                `json:"type,omitempty"`
	Probability *float32               `json:"probability,omitempty"`
}
