package bcc

type RenewElasticityAssuranceRequest struct {
	ElasticityAssuranceIds []*string `json:"elasticityAssuranceIds,omitempty"`
	Period                 *string   `json:"period,omitempty"`
}
