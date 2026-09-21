package vdb

type ListRecordsUsingGETRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	ListOrder  *string `json:"-"`
	Page       *string `json:"-"`
	PageSize   *string `json:"-"`
}
