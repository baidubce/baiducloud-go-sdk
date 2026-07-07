package agentidentity

type UpdateAgentRequest struct {
	AgentId                         *string   `json:"agentId,omitempty"`
	Description                     *string   `json:"description,omitempty"`
	AllowedResourceOauth2ReturnUrls []*string `json:"allowedResourceOauth2ReturnUrls,omitempty"`
}
