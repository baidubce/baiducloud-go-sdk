package rds

type ModifyAccountPermissionsRequest struct {
	InstanceId         *string              `json:"-"`
	AccountName        *string              `json:"-"`
	DatabasePrivileges []*DatabasePrivilege `json:"databasePrivileges,omitempty"`
}
