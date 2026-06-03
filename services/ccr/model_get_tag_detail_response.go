package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTagDetailResponse struct {
	bce.BaseResponse
	Architecture *string          `json:"architecture,omitempty"`
	Author       *string          `json:"author,omitempty"`
	Digest       *string          `json:"digest,omitempty"`
	Os           *string          `json:"os,omitempty"`
	ProjectId    *int32           `json:"projectId,omitempty"`
	PullTime     *string          `json:"pullTime,omitempty"`
	PushTime     *string          `json:"pushTime,omitempty"`
	RepositoryId *int32           `json:"repositoryId,omitempty"`
	ScanOverview *TagScanOverview `json:"scanOverview,omitempty"`
	Size         *int32           `json:"size,omitempty"`
	TagName      *string          `json:"tagName,omitempty"`
	CcrType      *string          `json:"type,omitempty"`
}
