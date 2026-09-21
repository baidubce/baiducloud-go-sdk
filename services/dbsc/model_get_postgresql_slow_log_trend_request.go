package dbsc

type GetPostgresqlSlowLogTrendRequest struct {
	Product *string `json:"-"`
	AppId   *string `json:"-"`
	NodeId  *string `json:"-"`
	Start   *string `json:"-"`
	End     *string `json:"-"`
	Period  *int32  `json:"-"`
}
