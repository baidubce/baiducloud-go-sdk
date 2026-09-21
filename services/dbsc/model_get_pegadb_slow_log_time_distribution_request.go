package dbsc

type GetPegadbSlowLogTimeDistributionRequest struct {
	AppId    *string `json:"-"`
	NodeId   *string `json:"-"`
	Start    *string `json:"-"`
	End      *string `json:"-"`
	DbEngine *string `json:"-"`
}
