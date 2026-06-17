package bcc

type CheckServerEventRequest struct {
	Action                        *string `json:"-"`
	ServerEventId                 *string `json:"serverEventId,omitempty"`
	InstanceId                    *string `json:"instanceId,omitempty"`
	CheckResult                   *string `json:"checkResult,omitempty"`
	IssueEffect                   *string `json:"issueEffect,omitempty"`
	IssueDescription              *string `json:"issueDescription,omitempty"`
	AuthorizeMaintenanceOperation *string `json:"authorizeMaintenanceOperation,omitempty"`
}
