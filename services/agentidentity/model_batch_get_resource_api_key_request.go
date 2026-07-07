package agentidentity

type BatchGetResourceApiKeyRequest struct {
	XBceWorkloadAccessToken *string   `json:"-"`
	Names                   []*string `json:"names,omitempty"`
	WorkloadAccessToken     *string   `json:"workloadAccessToken,omitempty"`
}
