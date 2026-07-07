package agentidentity

type GetResourceApikeyRequest struct {
	XBceWorkloadAccessToken *string `json:"-"`
	Name                    *string `json:"name,omitempty"`
	WorkloadAccessToken     *string `json:"workloadAccessToken,omitempty"`
}
