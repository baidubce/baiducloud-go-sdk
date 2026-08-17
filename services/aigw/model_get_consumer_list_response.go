package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetConsumerListResponse struct {
	bce.BaseResponse
	Success   *bool              `json:"success,omitempty"`
	Status    *int32             `json:"status,omitempty"`
	Total     *int32             `json:"total,omitempty"`
	Consumers []*ConsumerSummary `json:"consumers,omitempty"`
	NextToken *string            `json:"nextToken,omitempty"`
}
