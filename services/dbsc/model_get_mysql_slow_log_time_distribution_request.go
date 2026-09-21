package dbsc

type GetMysqlSlowLogTimeDistributionRequest struct {
	AppId          *string `json:"-"`
	NodeId         *string `json:"-"`
	Start          *string `json:"-"`
	End            *string `json:"-"`
	DbNames        *string `json:"-"`
	FingerprintMd5 *string `json:"-"`
}
