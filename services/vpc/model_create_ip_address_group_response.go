package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateIpAddressGroupResponse struct {
	bce.BaseResponse
	IpSetId *string `json:"ipSetId,omitempty"`
}
