package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RefreshRobotAccountKeyResponse struct {
	bce.BaseResponse
	Secret *string `json:"secret,omitempty"`
}
