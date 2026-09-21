package dbsc

type GetMongodbDatabaseSpaceTrendRequest struct {
	AppId      *string `json:"-"`
	NodeId     *string `json:"-"`
	Database   *string `json:"-"`
	Period     *int32  `json:"-"`
	Start      *string `json:"-"`
	End        *string `json:"-"`
	Metrics    *string `json:"-"`
	Statistics *string `json:"-"`
}
