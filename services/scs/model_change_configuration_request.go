package scs

type ChangeConfigurationRequest struct {
	InstanceId    *string  `json:"-"`
	ClientToken   *string  `json:"-"`
	Billing       *Billing `json:"billing,omitempty"`
	EngineVersion *string  `json:"engineVersion,omitempty"`
	NodeType      *string  `json:"nodeType,omitempty"`
	ShardNum      *int32   `json:"shardNum,omitempty"`
	DiskFlavor    *int32   `json:"diskFlavor,omitempty"`
}
