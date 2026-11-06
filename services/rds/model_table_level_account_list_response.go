package rds

type TableLevelAccountListResponse struct {
	User          *string `json:"user,omitempty"`
	Host          *string `json:"host,omitempty"`
	Status        *string `json:"status,omitempty"`
	RdsType       *string `json:"type,omitempty"`
	SuperUserFlag *string `json:"superUserFlag,omitempty"`
	Remark        *string `json:"remark,omitempty"`
	SupportBns    *bool   `json:"supportBns,omitempty"`
	WrFlag        *string `json:"wrFlag,omitempty"`
}
