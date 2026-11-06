package rds

type ErrorLogListRequest struct {
	InstanceId *string `json:"-"`
	Datetime   *string `json:"-"`
}
