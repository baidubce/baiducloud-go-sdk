package scs

type ModifyReplicationZoneRequest struct {
	InstanceId      *string            `json:"-"`
	IsDefer         *bool              `json:"isDefer,omitempty"`
	ReplicationInfo []*ReplicationItem `json:"replicationInfo,omitempty"`
}
