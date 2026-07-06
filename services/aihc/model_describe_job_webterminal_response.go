package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobWebterminalResponse struct {
	bce.BaseResponse
	RequestId      *string `json:"requestId,omitempty"`
	WebTerminalUrl *string `json:"webTerminalUrl,omitempty"`
}
