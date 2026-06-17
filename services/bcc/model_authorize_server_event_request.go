package bcc

type AuthorizeServerEventRequest struct {
	Action                        *string `json:"-"`
	ServerEventId                 *string `json:"serverEventId,omitempty"`
	InstanceId                    *string `json:"instanceId,omitempty"`
	AuthorizeMaintenanceOperation *string `json:"authorizeMaintenanceOperation,omitempty"`
	ExecuteTime                   *string `json:"executeTime,omitempty"`
}
