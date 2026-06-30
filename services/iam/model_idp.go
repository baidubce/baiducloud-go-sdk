package iam

type Idp struct {
	Status          *string `json:"status,omitempty"`
	DomainId        *string `json:"domainId,omitempty"`
	EncodeMetadata  *string `json:"encodeMetadata,omitempty"`
	FileName        *string `json:"fileName,omitempty"`
	AuxiliaryDomain *string `json:"auxiliaryDomain,omitempty"`
}
