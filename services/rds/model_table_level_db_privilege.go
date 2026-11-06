package rds

type TableLevelDbPrivilege struct {
	Db    *string   `json:"db,omitempty"`
	Privs []*string `json:"privs,omitempty"`
}
