package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryRoutingListResponse struct {
	bce.BaseResponse
	Result  *RouteListPage `json:"result,omitempty"`
	Message *string        `json:"message,omitempty"`
}
