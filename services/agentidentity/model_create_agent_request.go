package agentidentity

type CreateAgentRequest struct {
	Name                            *string   `json:"name,omitempty"`
	Description                     *string   `json:"description,omitempty"`
	AllowedResourceOauth2ReturnUrls []*string `json:"allowedResourceOauth2ReturnUrls,omitempty"`
}
