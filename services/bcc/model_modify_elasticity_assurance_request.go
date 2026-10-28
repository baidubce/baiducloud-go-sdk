package bcc

type ModifyElasticityAssuranceRequest struct {
	ElasticityAssuranceIds []*string `json:"elasticityAssuranceIds,omitempty"`
	AutoRenew              *bool     `json:"autoRenew,omitempty"`
	AutoRenewPeriod        *int32    `json:"autoRenewPeriod,omitempty"`
}
