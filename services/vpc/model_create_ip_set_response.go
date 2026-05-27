package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateIpSetResponse struct {
	bce.BaseResponse
	IpGroupId *string `json:"ipGroupId,omitempty"`
}
