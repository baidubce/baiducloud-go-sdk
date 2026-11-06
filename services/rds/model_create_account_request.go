package rds

type CreateAccountRequest struct {
	InstanceId         *string              `json:"-"`
	AccountName        *string              `json:"accountName,omitempty"`
	Password           *string              `json:"password,omitempty"`
	AccountType        *string              `json:"accountType,omitempty"`
	DatabasePrivileges []*DatabasePrivilege `json:"databasePrivileges,omitempty"`
	Desc               *string              `json:"desc,omitempty"`
	RdsType            *string              `json:"type,omitempty"`
}
