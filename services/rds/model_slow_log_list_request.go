package rds

type SlowLogListRequest struct {
	InstanceId *string `json:"-"`
	Datetime   *string `json:"-"`
}
