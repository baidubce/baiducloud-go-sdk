package rds

type ChangeAccountPasswordRequest struct {
	InstanceId  *string `json:"-"`
	AccountName *string `json:"-"`
	Password    *string `json:"password,omitempty"`
}
