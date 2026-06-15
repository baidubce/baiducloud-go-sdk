package bcc

type BatchDeleteIpRequest struct {
	InstanceId *string   `json:"instanceId,omitempty"`
	PrivateIps []*string `json:"privateIps,omitempty"`
}
