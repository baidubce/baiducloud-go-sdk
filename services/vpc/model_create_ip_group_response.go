package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateIpGroupResponse struct {
	bce.BaseResponse
	IpSetId *string `json:"ipSetId,omitempty"`
}
