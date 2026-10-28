package bcc

type QueryElasticityAssuranceRequest struct {
	Marker                  *string   `json:"marker,omitempty"`
	MaxKeys                 *int32    `json:"maxKeys,omitempty"`
	ElasticityAssuranceIds  []*string `json:"elasticityAssuranceIds,omitempty"`
	ElasticityAssuranceName *string   `json:"elasticityAssuranceName,omitempty"`
	ZoneName                *string   `json:"zoneName,omitempty"`
	Spec                    *string   `json:"spec,omitempty"`
	Status                  []*string `json:"status,omitempty"`
}
