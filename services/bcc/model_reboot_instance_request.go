package bcc

type RebootInstanceRequest struct {
	InstanceId *string `json:"-"`
	ForceStop  *bool   `json:"forceStop,omitempty"`
}
