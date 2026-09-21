package dbsc

type GetPostgresqlSlowLogTimeDistributionRequest struct {
	Product   *string   `json:"-"`
	AppId     *string   `json:"-"`
	NodeId    *string   `json:"-"`
	Start     *string   `json:"-"`
	End       *string   `json:"-"`
	DbNames   []*string `json:"-"`
	Users     []*string `json:"-"`
	ClientIPs []*string `json:"-"`
}
