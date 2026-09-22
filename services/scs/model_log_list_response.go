package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type LogListResponse struct {
	bce.BaseResponse
	LogList []*ShardLog `json:"logList,omitempty"`
}
