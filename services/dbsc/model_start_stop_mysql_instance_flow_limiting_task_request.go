package dbsc

type StartStopMysqlInstanceFlowLimitingTaskRequest struct {
	FilterId *int32  `json:"filterId,omitempty"`
	Action   *string `json:"action,omitempty"`
	AppId    *string `json:"appId,omitempty"`
	NodeId   *string `json:"nodeId,omitempty"`
}
