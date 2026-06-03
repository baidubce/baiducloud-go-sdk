package ccr

type UpdatePublicNetworkRequest struct {
	InstanceId *string `json:"-"`
	Action     *string `json:"action,omitempty"`
}
