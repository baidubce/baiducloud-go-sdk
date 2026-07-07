package agentidentity

type GetWATForUserRequest struct {
	BceUserId       *string `json:"bceUserId,omitempty"`
	AgentId         *string `json:"agentId,omitempty"`
	AgentName       *string `json:"agentName,omitempty"`
	UserId          *string `json:"userId,omitempty"`
	SessionId       *string `json:"sessionId,omitempty"`
	DurationSeconds *int32  `json:"durationSeconds,omitempty"`
}
