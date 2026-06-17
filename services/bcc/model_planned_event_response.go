package bcc

type PlannedEventResponse struct {
	ServerEventId                                *string                    `json:"serverEventId,omitempty"`
	ServerEventType                              *string                    `json:"serverEventType,omitempty"`
	ServerEventStatus                            *string                    `json:"serverEventStatus,omitempty"`
	InstanceId                                   *string                    `json:"instanceId,omitempty"`
	ProductCategory                              *string                    `json:"productCategory,omitempty"`
	InstanceSpec                                 *string                    `json:"instanceSpec,omitempty"`
	InstanceName                                 *string                    `json:"instanceName,omitempty"`
	PrivateIp                                    *string                    `json:"privateIp,omitempty"`
	Tags                                         []*TagModel                `json:"tags,omitempty"`
	ServerEventCreatedTime                       *string                    `json:"serverEventCreatedTime,omitempty"`
	ServerEventEndedTime                         *string                    `json:"serverEventEndedTime,omitempty"`
	MaintenanceOptions                           []*string                  `json:"maintenanceOptions,omitempty"`
	SupportMaintenanceOptions                    []*string                  `json:"supportMaintenanceOptions,omitempty"`
	AuthorizedMaintenanceOperation               *string                    `json:"authorizedMaintenanceOperation,omitempty"`
	AssociatedPlannedMaintenanceServerEventIds   []*string                  `json:"associatedPlannedMaintenanceServerEventIds,omitempty"`
	AssociatedUnplannedMaintenanceServerEventIds []*string                  `json:"associatedUnplannedMaintenanceServerEventIds,omitempty"`
	ExecuteTime                                  *string                    `json:"executeTime,omitempty"`
	ServerEventLogs                              []*OperationRecordResponse `json:"serverEventLogs,omitempty"`
	HasFastRepairStock                           *bool                      `json:"hasFastRepairStock,omitempty"`
	Risks                                        []*IssueResponse           `json:"risks,omitempty"`
}
