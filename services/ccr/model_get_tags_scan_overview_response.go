package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTagsScanOverviewResponse struct {
	bce.BaseResponse
	Items        []*ArtifactScanOverview `json:"items,omitempty"`
	LastScanTime *string                 `json:"lastScanTime,omitempty"`
	PageNo       *int32                  `json:"pageNo,omitempty"`
	PageSize     *int32                  `json:"pageSize,omitempty"`
	Summary      *map[string]interface{} `json:"summary,omitempty"`
	Total        *int32                  `json:"total,omitempty"`
}
