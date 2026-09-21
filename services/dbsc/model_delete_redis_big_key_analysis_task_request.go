package dbsc

type DeleteRedisBigKeyAnalysisTaskRequest struct {
	Ids   []*int64 `json:"-"`
	AppId *string  `json:"appId,omitempty"`
}
