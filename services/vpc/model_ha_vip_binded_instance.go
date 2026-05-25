package vpc

type HaVipBindedInstance struct {
	InstanceId   *string `json:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	Master       *bool   `json:"master,omitempty"`
}
