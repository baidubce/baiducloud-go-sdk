package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateGatewayLimitRulesResponse struct {
	bce.BaseResponse
	GlrId *string `json:"glrId,omitempty"`
}
