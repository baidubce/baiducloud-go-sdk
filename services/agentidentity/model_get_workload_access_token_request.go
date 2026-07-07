package agentidentity

type GetWorkloadAccessTokenRequest struct {
	BceUserId       *string `json:"bceUserId,omitempty"`
	AgentId         *string `json:"agentId,omitempty"`
	AgentName       *string `json:"agentName,omitempty"`
	DurationSeconds *int32  `json:"durationSeconds,omitempty"`
}
