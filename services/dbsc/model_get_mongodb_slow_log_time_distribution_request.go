package dbsc

type GetMongodbSlowLogTimeDistributionRequest struct {
	AppId          *string `json:"-"`
	NodeId         *string `json:"-"`
	Start          *string `json:"-"`
	End            *string `json:"-"`
	DbNames        *string `json:"-"`
	FingerprintMd5 *string `json:"-"`
}
