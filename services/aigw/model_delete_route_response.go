package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteRouteResponse struct {
	bce.BaseResponse
	Result  *DeleteRouteResult `json:"result,omitempty"`
	Message *string            `json:"message,omitempty"`
}
