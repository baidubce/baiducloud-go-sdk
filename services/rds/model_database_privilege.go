package rds

type DatabasePrivilege struct {
	DbName   *string `json:"dbName,omitempty"`
	AuthType *string `json:"authType,omitempty"`
}
