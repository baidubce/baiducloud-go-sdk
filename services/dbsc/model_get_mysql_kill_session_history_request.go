package dbsc

type GetMysqlKillSessionHistoryRequest struct {
	AppId  *string `json:"-"`
	NodeId *string `json:"-"`
	Start  *string `json:"-"`
	End    *string `json:"-"`
}
