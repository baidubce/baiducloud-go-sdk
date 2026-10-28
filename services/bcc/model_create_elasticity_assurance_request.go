package bcc

type CreateElasticityAssuranceRequest struct {
	ElasticityAssuranceName *string     `json:"elasticityAssuranceName,omitempty"`
	ZoneName                *string     `json:"zoneName,omitempty"`
	Spec                    *string     `json:"spec,omitempty"`
	Capacity                *int32      `json:"capacity,omitempty"`
	Period                  *int32      `json:"period,omitempty"`
	AutoRenew               *bool       `json:"autoRenew,omitempty"`
	Tags                    []*TagModel `json:"tags,omitempty"`
}
