package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeServiceConfigResponse struct {
	bce.BaseResponse
	Success                *bool                   `json:"success,omitempty"`
	Code                   *string                 `json:"code,omitempty"`
	Message                *string                 `json:"message,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	IncludeLLM             *bool                   `json:"includeLLM,omitempty"`
	ServiceDisplayName     *string                 `json:"serviceDisplayName,omitempty"`
	SampleConfig           *SampleConfig           `json:"sampleConfig,omitempty"`
	LoggingConfig          *LoggingConfig          `json:"loggingConfig,omitempty"`
	RequestConfig          *ServiceRequestConfig   `json:"requestConfig,omitempty"`
	TopoConfig             *ServiceTopoConfig      `json:"topoConfig,omitempty"`
	MllmResourceDumpConfig *MllmResourceDumpConfig `json:"mllmResourceDumpConfig,omitempty"`
}
