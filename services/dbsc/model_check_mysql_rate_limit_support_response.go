package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CheckMysqlRateLimitSupportResponse struct {
	bce.BaseResponse
	Allowed *bool `json:"allowed,omitempty"`
}
