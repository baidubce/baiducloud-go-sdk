package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RemoteReadResponse struct {
	bce.BaseResponse
	Status           *string                   `json:"status,omitempty"`
	IsPartial        *bool                     `json:"isPartial,omitempty"`
	Data             *map[string]interface{}   `json:"data,omitempty"`
	DataResultType   *string                   `json:"data.resultType,omitempty"`
	DataResult       []*map[string]interface{} `json:"data.result,omitempty"`
	DataResultMetric *map[string]interface{}   `json:"data.result[].metric,omitempty"`
	DataResultValues []*map[string]interface{} `json:"data.result[].values,omitempty"`
	DataResultValue  []*map[string]interface{} `json:"data.result[].value,omitempty"`
}
