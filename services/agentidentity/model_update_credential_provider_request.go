package agentidentity

type UpdateCredentialProviderRequest struct {
	CredentialProviderId *string      `json:"credentialProviderId,omitempty"`
	Desc                 *string      `json:"desc,omitempty"`
	Credential           *interface{} `json:"credential,omitempty"`
}
