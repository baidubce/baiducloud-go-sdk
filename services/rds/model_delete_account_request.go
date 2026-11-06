package rds

type DeleteAccountRequest struct {
	InstanceId  *string `json:"-"`
	AccountName *string `json:"-"`
}
