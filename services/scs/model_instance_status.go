package scs

// InstanceStatus the model 'InstanceStatus'
type InstanceStatus string

// List of InstanceStatus
const (
	InstanceStatusValueUnknown                InstanceStatus = "状态值"
	InstanceStatusCreating                    InstanceStatus = "Creating"
	InstanceStatusDataprepareing              InstanceStatus = "Dataprepareing"
	InstanceStatusRunning                     InstanceStatus = "Running"
	InstanceStatusValueToSync                 InstanceStatus = "To-sync"
	InstanceStatusRebooting                   InstanceStatus = "Rebooting"
	InstanceStatusPausing                     InstanceStatus = "Pausing"
	InstanceStatusPaused                      InstanceStatus = "Paused"
	InstanceStatusDeleted                     InstanceStatus = "Deleted"
	InstanceStatusDeleting                    InstanceStatus = "Deleting"
	InstanceStatusError                       InstanceStatus = "Error"
	InstanceStatusFailed                      InstanceStatus = "Failed"
	InstanceStatusModifying                   InstanceStatus = "Modifying"
	InstanceStatusModifyfailed                InstanceStatus = "Modifyfailed"
	InstanceStatusExpired                     InstanceStatus = "Expired"
	InstanceStatusFlushing                    InstanceStatus = "Flushing"
	InstanceStatusValueFlushFailed            InstanceStatus = "Flush failed"
	InstanceStatusAztransforming              InstanceStatus = "Aztransforming"
	InstanceStatusBackuping                   InstanceStatus = "Backuping"
	InstanceStatusRecovering                  InstanceStatus = "Recovering"
	InstanceStatusRestarting                  InstanceStatus = "Restarting"
	InstanceStatusAnalysising                 InstanceStatus = "Analysising"
	InstanceStatusExchanging                  InstanceStatus = "Exchanging"
	InstanceStatusValueExchangeFailed         InstanceStatus = "Exchange failed"
	InstanceStatusIsolated                    InstanceStatus = "Isolated"
	InstanceStatusRelaunchable                InstanceStatus = "Relaunchable"
	InstanceStatusModifyable                  InstanceStatus = "Modifyable"
	InstanceStatusTopomodifing                InstanceStatus = "Topomodifing"
	InstanceStatusQuitting                    InstanceStatus = "Quitting"
	InstanceStatusRoCreating                  InstanceStatus = "Ro_creating"
	InstanceStatusRoDeleting                  InstanceStatus = "Ro_deleting"
	InstanceStatusUpdateDomainInDb            InstanceStatus = "Update_domain_in_db"
	InstanceStatusSynccreating                InstanceStatus = "Synccreating"
	InstanceStatusSyncdeleting                InstanceStatus = "Syncdeleting"
	InstanceStatusEntranceCreating            InstanceStatus = "Entrance_creating"
	InstanceStatusModifyingAz                 InstanceStatus = "Modifying_az"
	InstanceStatusModifyingDefaultEntrance    InstanceStatus = "Modifying_default_entrance"
	InstanceStatusModifyingNetBandwidth       InstanceStatus = "Modifying_net_bandwidth"
	InstanceStatusNormal                      InstanceStatus = "Normal"
	InstanceStatusInitial                     InstanceStatus = "Initial"
	InstanceStatusValueDelMember              InstanceStatus = "Del-member"
	InstanceStatusGroupmodifying              InstanceStatus = "Groupmodifying"
	InstanceStatusValueAzTransformFailed      InstanceStatus = "Az transform failed"
	InstanceStatusValueModifyTypeAddShards    InstanceStatus = "Modify-type-add-shards"
	InstanceStatusValueModifyTypeDelShards    InstanceStatus = "Modify-type-del-shards"
	InstanceStatusValueModifyTypeIncrNodeType InstanceStatus = "Modify-type-incr-node-type"
	InstanceStatusValueModifyTypeDecrNodeType InstanceStatus = "Modify-type-decr-node-type"
	InstanceStatusValueModifyTypeSpec         InstanceStatus = "Modify-type-spec"
	InstanceStatusConfiguring                 InstanceStatus = "Configuring"
	InstanceStatusValueAddMember              InstanceStatus = "Add-member"
	InstanceStatusMalfunctioning              InstanceStatus = "Malfunctioning"
	InstanceStatusProxyReplacing              InstanceStatus = "Proxy_replacing"
)

// All allowed values of InstanceStatus enum
var AllowedInstanceStatusEnumValues = []InstanceStatus{
	"状态值",
	"Creating",
	"Dataprepareing",
	"Running",
	"To-sync",
	"Rebooting",
	"Pausing",
	"Paused",
	"Deleted",
	"Deleting",
	"Error",
	"Failed",
	"Modifying",
	"Modifyfailed",
	"Expired",
	"Flushing",
	"Flush failed",
	"Aztransforming",
	"Backuping",
	"Recovering",
	"Restarting",
	"Analysising",
	"Exchanging",
	"Exchange failed",
	"Isolated",
	"Relaunchable",
	"Modifyable",
	"Topomodifing",
	"Quitting",
	"Ro_creating",
	"Ro_deleting",
	"Update_domain_in_db",
	"Synccreating",
	"Syncdeleting",
	"Entrance_creating",
	"Modifying_az",
	"Modifying_default_entrance",
	"Modifying_net_bandwidth",
	"Normal",
	"Initial",
	"Del-member",
	"Groupmodifying",
	"Az transform failed",
	"Modify-type-add-shards",
	"Modify-type-del-shards",
	"Modify-type-incr-node-type",
	"Modify-type-decr-node-type",
	"Modify-type-spec",
	"Configuring",
	"Add-member",
	"Malfunctioning",
	"Proxy_replacing",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceStatus) IsValid() bool {
	for _, existing := range AllowedInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
