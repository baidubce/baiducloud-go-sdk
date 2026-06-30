package iam

type UpdateSubUserIdpRequest struct {
	FileName        *string `json:"fileName,omitempty"`
	EncodeMetadata  *string `json:"encodeMetadata,omitempty"`
	AuxiliaryDomain *string `json:"auxiliaryDomain,omitempty"`
}
