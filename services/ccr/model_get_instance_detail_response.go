package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceDetailResponse struct {
	bce.BaseResponse
	Bucket    *string            `json:"bucket,omitempty"`
	Region    *string            `json:"region,omitempty"`
	Info      *Instance          `json:"info,omitempty"`
	Statistic *InstanceStatistic `json:"statistic,omitempty"`
	Quota     *InstanceQuota     `json:"quota,omitempty"`
}
