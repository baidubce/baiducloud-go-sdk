package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RetrieveDatabaseListResponse struct {
	bce.BaseResponse
	Databases []*Database `json:"databases,omitempty"`
}
