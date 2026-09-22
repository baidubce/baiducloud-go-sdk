package scs

type GePriceForResizeInstanceRequest struct {
	ClientToken    *string `json:"-"`
	InstanceId     *string `json:"-"`
	NodeType       *string `json:"nodeType,omitempty"`
	ShardNum       *int32  `json:"shardNum,omitempty"`
	ReplicationNum *int32  `json:"replicationNum,omitempty"`
	DiskFlavor     *int32  `json:"diskFlavor,omitempty"`
	ChargeType     *string `json:"chargeType,omitempty"`
	Period         *int32  `json:"period,omitempty"`
	ChangeType     *string `json:"changeType,omitempty"`
}
