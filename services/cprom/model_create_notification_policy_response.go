package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateNotificationPolicyResponse struct {
	bce.BaseResponse
	NotifyRuleId *string `json:"notifyRuleId,omitempty"`
}
