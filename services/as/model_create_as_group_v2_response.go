package as

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAsGroupV2Response struct {
	bce.BaseResponse
	GroupId *string   `json:"groupId,omitempty"`
	OrderId []*string `json:"orderId,omitempty"`
}
