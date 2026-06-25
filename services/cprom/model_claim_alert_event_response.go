package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ClaimAlertEventResponse struct {
	bce.BaseResponse
	SuccessCount *int32              `json:"successCount,omitempty"`
	FailedCount  *int32              `json:"failedCount,omitempty"`
	Details      []*EventClaimDetail `json:"details,omitempty"`
}
