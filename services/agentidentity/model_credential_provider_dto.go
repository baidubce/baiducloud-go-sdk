package agentidentity

type CredentialProviderDTO struct {
	Id                *string           `json:"id,omitempty"`
	BceDomainId       *string           `json:"bceDomainId,omitempty"`
	BceUserId         *string           `json:"bceUserId,omitempty"`
	Name              *string           `json:"name,omitempty"`
	AgentidentityType *string           `json:"type,omitempty"`
	Desc              *string           `json:"desc,omitempty"`
	Credential        *CredentialConfig `json:"credential,omitempty"`
	CreatedAt         *string           `json:"createdAt,omitempty"`
	UpdatedAt         *string           `json:"updatedAt,omitempty"`
}
