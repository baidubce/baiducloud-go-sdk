package rds

type CreateADatabaseRequest struct {
	InstanceId        *string             `json:"-"`
	CharacterSetName  *string             `json:"characterSetName,omitempty"`
	DbName            *string             `json:"dbName,omitempty"`
	Remark            *string             `json:"remark,omitempty"`
	AccountPrivileges []*AccountPrivilege `json:"accountPrivileges,omitempty"`
	Ctype             *string             `json:"ctype,omitempty"`
	Collate           *string             `json:"collate,omitempty"`
	Owner             *string             `json:"owner,omitempty"`
}
