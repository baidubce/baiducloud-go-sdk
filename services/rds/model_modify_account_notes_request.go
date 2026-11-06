package rds

type ModifyAccountNotesRequest struct {
	InstanceId  *string `json:"-"`
	AccountName *string `json:"-"`
	Remark      *string `json:"remark,omitempty"`
}
