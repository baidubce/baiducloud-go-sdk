package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateRobotAccountResponse struct {
	bce.BaseResponse
	Id           *int32  `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Secret       *string `json:"secret,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
	ExpiresAt    *int32  `json:"expiresAt,omitempty"`
}
