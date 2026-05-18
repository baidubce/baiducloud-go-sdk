package eip

type CreateEipBpRequest struct {
	ClientToken     *string     `json:"-"`
	Name            *string     `json:"name,omitempty"`
	Eip             *string     `json:"eip,omitempty"`
	EipGroupId      *string     `json:"eipGroupId,omitempty"`
	BandwidthInMbps *int32      `json:"bandwidthInMbps,omitempty"`
	EipType         *string     `json:"type,omitempty"`
	AutoReleaseTime *string     `json:"autoReleaseTime,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	ResourceGroupId *string     `json:"resourceGroupId,omitempty"`
}
