package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TableGranularityQueryUserDetailsApiResponse struct {
	bce.BaseResponse
	User          *string                     `json:"user,omitempty"`
	Host          *string                     `json:"host,omitempty"`
	Password      *string                     `json:"password,omitempty"`
	RdsType       *string                     `json:"type,omitempty"`
	SuperUserFlag *string                     `json:"superUserFlag,omitempty"`
	Remark        *string                     `json:"remark,omitempty"`
	Privilege     *TableLevelPrivilegeRequest `json:"privilege,omitempty"`
}
