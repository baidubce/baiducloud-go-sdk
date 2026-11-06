package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryAccountListResponse struct {
	bce.BaseResponse
	Accounts []*Account `json:"accounts,omitempty"`
}
