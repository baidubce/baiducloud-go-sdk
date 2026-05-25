package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateIpAddressFamilyResponse struct {
	bce.BaseResponse
	IpGroupId *string `json:"ipGroupId,omitempty"`
}
