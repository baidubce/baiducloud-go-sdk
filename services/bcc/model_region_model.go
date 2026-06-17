package bcc

type RegionModel struct {
	RegionId       *string `json:"regionId,omitempty"`
	RegionName     *string `json:"regionName,omitempty"`
	RegionEndpoint *string `json:"regionEndpoint,omitempty"`
}
