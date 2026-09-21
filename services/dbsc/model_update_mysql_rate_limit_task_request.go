package dbsc

type UpdateMysqlRateLimitTaskRequest struct {
	FilterId    *int32  `json:"filterId,omitempty"`
	AppId       *string `json:"appId,omitempty"`
	NodeId      *string `json:"nodeId,omitempty"`
	FilterKey   *string `json:"filterKey,omitempty"`
	FilterLimit *int32  `json:"filterLimit,omitempty"`
	FilterType  *string `json:"filterType,omitempty"`
}
