package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDefaultConfigResponse struct {
	bce.BaseResponse
	Success       *bool          `json:"success,omitempty"`
	Code          *string        `json:"code,omitempty"`
	Message       *string        `json:"message,omitempty"`
	StorageConfig *StorageConfig `json:"storageConfig,omitempty"`
	RequestConfig *RequestConfig `json:"requestConfig,omitempty"`
	TopoConfig    *TopoConfig    `json:"topoConfig,omitempty"`
}
