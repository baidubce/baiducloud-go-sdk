package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbSlowQueryTemplateResponse struct {
	bce.BaseResponse
	Items      []*MongoDBSlowLogTemplate `json:"items,omitempty"`
	TotalCount *int32                    `json:"totalCount,omitempty"`
}
