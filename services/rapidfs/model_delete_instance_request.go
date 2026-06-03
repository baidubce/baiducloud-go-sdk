package rapidfs

type DeleteInstanceRequest struct {
	Action      *string `json:"-"`
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	Token       *string `json:"token,omitempty"`
}
