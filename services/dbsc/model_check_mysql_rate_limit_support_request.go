package dbsc

type CheckMysqlRateLimitSupportRequest struct {
	AppId  *string `json:"-"`
	NodeId *string `json:"-"`
}
