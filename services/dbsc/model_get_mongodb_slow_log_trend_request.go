package dbsc

type GetMongodbSlowLogTrendRequest struct {
	AppId  *string `json:"-"`
	NodeId *string `json:"-"`
	Start  *string `json:"-"`
	End    *string `json:"-"`
	Period *string `json:"-"`
}
