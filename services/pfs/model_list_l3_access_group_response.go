package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListL3AccessGroupResponse struct {
	bce.BaseResponse
	RequestId      *string            `json:"requestId,omitempty"`
	Result         []*AccessGroupInfo `json:"result,omitempty"`
	PageNo         *int32             `json:"pageNo,omitempty"`
	PageSize       *int32             `json:"pageSize,omitempty"`
	TotalCount     *int32             `json:"totalCount,omitempty"`
	RealTotalCount *int32             `json:"realTotalCount,omitempty"`
	Order          *string            `json:"order,omitempty"`
	OrderBy        *string            `json:"orderBy,omitempty"`
}
