package rds

type Account struct {
	AccountName        *string              `json:"accountName,omitempty"`
	DatabasePrivileges []*DatabasePrivilege `json:"databasePrivileges,omitempty"`
	RdsType            *string              `json:"type,omitempty"`
	AccountStatus      *string              `json:"accountStatus,omitempty"`
	SuperUserFlag      *string              `json:"superUserFlag,omitempty"`
	AccountType        *string              `json:"accountType,omitempty"`
	Desc               *string              `json:"desc,omitempty"`
	Status             *string              `json:"status,omitempty"`
	ETag               *string              `json:"ETag,omitempty"`
}
