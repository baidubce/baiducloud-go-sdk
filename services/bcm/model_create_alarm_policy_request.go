package bcm

type CreateAlarmPolicyRequest struct {
	Name                       *string         `json:"name,omitempty"`
	Scope                      *string         `json:"scope,omitempty"`
	ResourceType               *string         `json:"resourceType,omitempty"`
	Target                     *AlarmTarget    `json:"target,omitempty"`
	Rules                      []*AlarmRule    `json:"rules,omitempty"`
	PendingCount               *int32          `json:"pendingCount,omitempty"`
	OnMissingData              *string         `json:"onMissingData,omitempty"`
	NoDataNotifyPendingMinutes *int32          `json:"noDataNotifyPendingMinutes,omitempty"`
	BcmType                    *string         `json:"type,omitempty"`
	Level                      *string         `json:"level,omitempty"`
	Actions                    []*PolicyAction `json:"actions,omitempty"`
	NotifyEnabled              *bool           `json:"notifyEnabled,omitempty"`
	Callbacks                  []*Callback     `json:"callbacks,omitempty"`
	RenotifyCount              *int32          `json:"renotifyCount,omitempty"`
	RenotifyIntervalMinutes    *int32          `json:"renotifyIntervalMinutes,omitempty"`
	NotifyMergeWindowSeconds   *int32          `json:"notifyMergeWindowSeconds,omitempty"`
}
