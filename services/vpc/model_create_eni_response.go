package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEniResponse struct {
	bce.BaseResponse
	EniId *string `json:"eniId,omitempty"`
}
