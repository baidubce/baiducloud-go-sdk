package dbsc

type KillMysqlSessionRequest struct {
	AppId   *string  `json:"appId,omitempty"`
	NodeId  *string  `json:"nodeId,omitempty"`
	IdItems []*int32 `json:"idItems,omitempty"`
}
