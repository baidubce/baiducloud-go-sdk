package dbsc

type GetMysqlRateLimitTaskDetailRequest struct {
	AppId    *string `json:"-"`
	NodeId   *string `json:"-"`
	FilterId *int32  `json:"-"`
}
