package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteServiceResponse struct {
	bce.BaseResponse
	Result *string `json:"result,omitempty"`
}
