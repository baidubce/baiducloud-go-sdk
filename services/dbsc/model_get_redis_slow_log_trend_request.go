package dbsc

type GetRedisSlowLogTrendRequest struct {
	AppId    *string `json:"-"`
	NodeId   *string `json:"-"`
	Start    *string `json:"-"`
	End      *string `json:"-"`
	DbEngine *string `json:"-"`
	Period   *int32  `json:"-"`
}
