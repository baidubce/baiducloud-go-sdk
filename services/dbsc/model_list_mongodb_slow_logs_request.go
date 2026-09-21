package dbsc

type ListMongodbSlowLogsRequest struct {
	AppId          *string `json:"-"`
	NodeId         *string `json:"-"`
	Start          *string `json:"-"`
	End            *string `json:"-"`
	Users          *string `json:"-"`
	DbNames        *string `json:"-"`
	ClientIps      *string `json:"-"`
	Namespace      *string `json:"-"`
	FingerprintMd5 *string `json:"-"`
	OrderBy        *string `json:"-"`
	Order          *string `json:"-"`
	Page           *int32  `json:"-"`
	PageSize       *int32  `json:"-"`
}
