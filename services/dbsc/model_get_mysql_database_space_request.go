package dbsc

type GetMysqlDatabaseSpaceRequest struct {
	AppId    *string `json:"-"`
	NodeId   *string `json:"-"`
	Database *string `json:"-"`
	OrderBy  *string `json:"-"`
	Order    *string `json:"-"`
	Page     *int32  `json:"-"`
	PageSize *int32  `json:"-"`
}
