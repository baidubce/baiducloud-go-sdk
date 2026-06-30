package bci

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateImageCacheResponse struct {
	bce.BaseResponse
	ImageCacheId *string `json:"imageCacheId,omitempty"`
}
