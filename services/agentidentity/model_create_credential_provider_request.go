package agentidentity

type CreateCredentialProviderRequest struct {
	Name              *string      `json:"name,omitempty"`
	AgentidentityType *string      `json:"type,omitempty"`
	Desc              *string      `json:"desc,omitempty"`
	Credential        *interface{} `json:"credential,omitempty"`
}
