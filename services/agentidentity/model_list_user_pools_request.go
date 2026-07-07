package agentidentity

type ListUserPoolsRequest struct {
	Keyword  *string `json:"keyword,omitempty"`
	PageNo   *int32  `json:"pageNo,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty"`
}
