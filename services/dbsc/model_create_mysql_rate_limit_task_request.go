package dbsc

type CreateMysqlRateLimitTaskRequest struct {
	AppId       *string `json:"appId,omitempty"`
	NodeId      *string `json:"nodeId,omitempty"`
	FilterKey   *string `json:"filterKey,omitempty"`
	FilterLimit *int32  `json:"filterLimit,omitempty"`
	FilterType  *string `json:"filterType,omitempty"`
}
