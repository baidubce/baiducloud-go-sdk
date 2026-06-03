package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAuthGroupResponse struct {
	bce.BaseResponse
	AuthGroupInfo *AuthGroupInfo `json:"authGroupInfo,omitempty"`
}
