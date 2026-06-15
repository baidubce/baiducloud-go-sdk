package bcc

type EnterRescueModeRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	ForceStop  *bool   `json:"forceStop,omitempty"`
	Password   *string `json:"password,omitempty"`
}
