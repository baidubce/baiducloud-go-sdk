package rds

type QueryAccountDetailsRequest struct {
	InstanceId  *string `json:"-"`
	AccountName *string `json:"-"`
}
