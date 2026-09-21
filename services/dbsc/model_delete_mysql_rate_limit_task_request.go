package dbsc

type DeleteMysqlRateLimitTaskRequest struct {
	FilterId *int32  `json:"-"`
	AppId    *string `json:"-"`
	NodeId   *string `json:"-"`
}
