package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeTokenResponse struct {
	bce.BaseResponse
	TokenInfo *TokenInfo `json:"tokenInfo,omitempty"`
}
