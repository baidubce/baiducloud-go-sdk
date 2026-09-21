package dbsc

type GetMongodbCollectionSpaceTrendRequest struct {
	AppId      *string `json:"-"`
	NodeId     *string `json:"-"`
	Database   *string `json:"-"`
	Collection *string `json:"-"`
	Period     *int32  `json:"-"`
	Start      *string `json:"-"`
	End        *string `json:"-"`
	Metrics    *string `json:"-"`
	Statistics *string `json:"-"`
}
