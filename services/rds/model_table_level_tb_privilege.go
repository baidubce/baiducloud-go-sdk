package rds

type TableLevelTbPrivilege struct {
	Db    *string   `json:"db,omitempty"`
	Table *string   `json:"table,omitempty"`
	Privs []*string `json:"privs,omitempty"`
}
