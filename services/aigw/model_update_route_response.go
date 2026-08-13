package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateRouteResponse struct {
	bce.BaseResponse
	Result  *RouteResult `json:"result,omitempty"`
	Message *string      `json:"message,omitempty"`
}
