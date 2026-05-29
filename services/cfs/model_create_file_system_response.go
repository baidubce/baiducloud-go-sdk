package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateFileSystemResponse struct {
	bce.BaseResponse
	FsId *string `json:"fsId,omitempty"`
}
