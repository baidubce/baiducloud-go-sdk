package dbsc

type ListPostgresqlSlowLogsRequest struct {
	AppId     *string   `json:"-"`
	NodeId    *string   `json:"-"`
	Start     *string   `json:"-"`
	End       *string   `json:"-"`
	Page      *int32    `json:"-"`
	PageSize  *int32    `json:"-"`
	DbNames   []*string `json:"-"`
	ClientIPs []*string `json:"-"`
	Users     []*string `json:"-"`
	OrderBy   *string   `json:"-"`
	Order     *string   `json:"-"`
}
