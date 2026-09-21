package dbsc

type ListPegadbSlowLogsRequest struct {
	AppId    *string `json:"-"`
	NodeId   *string `json:"-"`
	Start    *string `json:"-"`
	End      *string `json:"-"`
	DbEngine *string `json:"-"`
	Page     *int32  `json:"-"`
	PageSize *int32  `json:"-"`
	OrderBy  *string `json:"-"`
	Order    *string `json:"-"`
}
