package agentidentity

type ListAgentsRequest struct {
	PageNo   *int32  `json:"pageNo,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty"`
	Keyword  *string `json:"keyword,omitempty"`
}
