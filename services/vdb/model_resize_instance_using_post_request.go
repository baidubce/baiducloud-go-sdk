package vdb

type ResizeInstanceUsingPOSTRequest struct {
	EngineType     *string            `json:"-"`
	Components     []*MilvusComponent `json:"components,omitempty"`
	DataNodeNum    *int32             `json:"dataNodeNum,omitempty"`
	DiskFlavor     *int32             `json:"diskFlavor,omitempty"`
	DiskType       *string            `json:"diskType,omitempty"`
	Env            *string            `json:"env,omitempty"`
	InstanceId     *string            `json:"instanceId,omitempty"`
	MasterNodeSpec *string            `json:"masterNodeSpec,omitempty"`
	MasterNum      *int32             `json:"masterNum,omitempty"`
	NodeSpec       *string            `json:"nodeSpec,omitempty"`
	NodeType       *string            `json:"nodeType,omitempty"`
	OrderId        *string            `json:"orderId,omitempty"`
	ProxyNodeSpec  *string            `json:"proxyNodeSpec,omitempty"`
	ProxyNum       *int32             `json:"proxyNum,omitempty"`
}
