package rapidfs

type DeleteInstanceRequest struct {
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	Token       *string `json:"token,omitempty"`
}
