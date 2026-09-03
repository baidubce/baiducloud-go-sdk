package aigw

type UpdateAIGatewayRequest struct {
	InstanceId       *string   `json:"-"`
	XRegion          *string   `json:"-"`
	Name             *string   `json:"name,omitempty"`
	Description      *string   `json:"description,omitempty"`
	DeleteProtection *bool     `json:"deleteProtection,omitempty"`
	PublicAccessible *bool     `json:"publicAccessible,omitempty"`
	Replicas         *int32    `json:"replicas,omitempty"`
	NetworkTypes     []*string `json:"networkTypes,omitempty"`
	Tags             []*Tag    `json:"tags,omitempty"`
}
