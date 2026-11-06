package rds

type TableLevelPrivilegeRequest struct {
	RdsType   *string                  `json:"type,omitempty"`
	Globals   []*string                `json:"globals,omitempty"`
	Databases []*TableLevelDbPrivilege `json:"databases,omitempty"`
	Tables    []*TableLevelTbPrivilege `json:"tables,omitempty"`
	Proxys    []*ProxyWrFlagInfo       `json:"proxys,omitempty"`
}
