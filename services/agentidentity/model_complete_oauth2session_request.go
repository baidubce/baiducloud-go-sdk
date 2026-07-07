package agentidentity

type CompleteOauth2sessionRequest struct {
	XBceWorkloadAccessToken *string      `json:"-"`
	SessionUri              *string      `json:"sessionUri,omitempty"`
	UserIdentifier          *interface{} `json:"userIdentifier,omitempty"`
	UserIdentifierUserId    *string      `json:"userIdentifier.userId,omitempty"`
	WorkloadAccessToken     *string      `json:"workloadAccessToken,omitempty"`
}
