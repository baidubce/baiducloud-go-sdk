package dbsc

type ListMysqlSlowLogsRequest struct {
	AppId          *string `json:"-"`
	NodeId         *string `json:"-"`
	Start          *string `json:"-"`
	End            *string `json:"-"`
	Users          *string `json:"-"`
	DbNames        *string `json:"-"`
	ClientIps      *string `json:"-"`
	FingerprintMd5 *string `json:"-"`
	OrderBy        *string `json:"-"`
	Order          *string `json:"-"`
	Page           *int32  `json:"-"`
	PageSize       *int32  `json:"-"`
}
