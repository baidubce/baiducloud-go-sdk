package et

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryDedicatedLineDetailResponse struct {
	bce.BaseResponse
	Id          *string     `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *string     `json:"status,omitempty"`
	ExpireTime  *string     `json:"expireTime,omitempty"`
	Isp         *string     `json:"isp,omitempty"`
	IntfType    *string     `json:"intfType,omitempty"`
	ApType      *string     `json:"apType,omitempty"`
	ApAddr      *string     `json:"apAddr,omitempty"`
	LinkDelay   *int32      `json:"linkDelay,omitempty"`
	UserName    *string     `json:"userName,omitempty"`
	UserPhone   *string     `json:"userPhone,omitempty"`
	UserEmail   *string     `json:"userEmail,omitempty"`
	UserIdc     *string     `json:"userIdc,omitempty"`
	Tags        []*TagModel `json:"tags,omitempty"`
}
