package agentidentity

type ListCredentialProvidersRequest struct {
	PageNo            *int32  `json:"pageNo,omitempty"`
	PageSize          *int32  `json:"pageSize,omitempty"`
	AgentidentityType *string `json:"type,omitempty"`
	Name              *string `json:"name,omitempty"`
}
