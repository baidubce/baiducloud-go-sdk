package aigw

type DeleteAIGatewayRequest struct {
	InstanceId *string `json:"-"`
	Force      *bool   `json:"-"`
	XRegion    *string `json:"-"`
}
