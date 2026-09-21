package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type KillMysqlSessionResponse struct {
	bce.BaseResponse
	Success *bool `json:"success,omitempty"`
}
