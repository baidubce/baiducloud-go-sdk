package scs

type DomainNameExchangeRequest struct {
	SourceInstanceId *string `json:"sourceInstanceId,omitempty"`
	TargetInstanceId *string `json:"targetInstanceId,omitempty"`
}
