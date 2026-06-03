package rapidfs

type ResizeInstanceRequest struct {
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	CapacityTiB *int32  `json:"capacityTiB,omitempty"`
}
