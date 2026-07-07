package agentidentity

type ListOauth2ClientsRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Keyword    *string `json:"keyword,omitempty"`
	PageNo     *int32  `json:"pageNo,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty"`
}
