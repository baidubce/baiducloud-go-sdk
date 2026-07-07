package agentidentity

type ResourceCredentialDTO struct {
	AgentidentityType *string      `json:"type,omitempty"`
	Name              *string      `json:"name,omitempty"`
	Credential        *interface{} `json:"credential,omitempty"`
	CredentialApiKey  *string      `json:"credential.apiKey,omitempty"`
	ExpireAt          *string      `json:"expireAt,omitempty"`
}
