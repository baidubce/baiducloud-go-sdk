package rds

type Database struct {
	DbName            *string             `json:"dbName,omitempty"`
	CharacterSetName  *string             `json:"characterSetName,omitempty"`
	AccountPrivileges []*AccountPrivilege `json:"accountPrivileges,omitempty"`
	Collate           *string             `json:"collate,omitempty"`
	Ctype             *string             `json:"ctype,omitempty"`
	Owner             *string             `json:"owner,omitempty"`
}
