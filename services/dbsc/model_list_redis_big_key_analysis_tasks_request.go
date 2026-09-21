package dbsc

type ListRedisBigKeyAnalysisTasksRequest struct {
	Id        *int32  `json:"-"`
	AppId     *string `json:"-"`
	ClusterId *string `json:"-"`
	DataType  *string `json:"-"`
	Order     *string `json:"-"`
	OrderBy   *string `json:"-"`
}
