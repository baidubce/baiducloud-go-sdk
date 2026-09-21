package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbSpaceSummaryResponse struct {
	bce.BaseResponse
	TotalDiskSize   *string `json:"totalDiskSize,omitempty"`
	UsedSpace       *string `json:"usedSpace,omitempty"`
	AvailableSpace  *string `json:"availableSpace,omitempty"`
	AvailableDays   *string `json:"availableDays,omitempty"`
	DayGrowSpaceAvg *string `json:"dayGrowSpaceAvg,omitempty"`
}
