package dbsc

type GetMysqlTableIndexesRequest struct {
	Product  *string `json:"-"`
	AppId    *string `json:"-"`
	Database *string `json:"-"`
	Table    *string `json:"-"`
}
