package agentidentity

type BatchAcquisitionOfUsersRequest struct {
	UserPoolId *string   `json:"userPoolId,omitempty"`
	Ids        []*string `json:"ids,omitempty"`
}
