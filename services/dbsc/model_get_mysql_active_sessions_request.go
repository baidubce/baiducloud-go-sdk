package dbsc

type GetMysqlActiveSessionsRequest struct {
	AppId  *string `json:"-"`
	NodeId *string `json:"-"`
}
