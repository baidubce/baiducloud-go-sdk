package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetAlertDetailResponse struct {
	bce.BaseResponse
	AlertId      *string            `json:"alertId,omitempty"`
	AlertName    *string            `json:"alertName,omitempty"`
	Expr         *string            `json:"expr,omitempty"`
	CpromFor     *string            `json:"for,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Enable       *bool              `json:"enable,omitempty"`
	NotifyRuleId *string            `json:"notifyRuleId,omitempty"`
	Severity     *string            `json:"severity,omitempty"`
	Annotations  *map[string]string `json:"annotations,omitempty"`
}
