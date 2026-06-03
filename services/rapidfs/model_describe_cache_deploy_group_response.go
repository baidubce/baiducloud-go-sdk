package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheDeployGroupResponse struct {
	bce.BaseResponse
	CacheDeployGroupInfo *CacheDeployGroupInfo `json:"cacheDeployGroupInfo,omitempty"`
}
