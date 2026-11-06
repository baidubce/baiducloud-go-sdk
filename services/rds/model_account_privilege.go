package rds

type AccountPrivilege struct {
	AccountName *string `json:"accountName,omitempty"`
	AuthType    *string `json:"authType,omitempty"`
}
