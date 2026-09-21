package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbDatabaseSpaceResponse struct {
	bce.BaseResponse
	Items          []*MongodbDatabaseSpaceItem `json:"items,omitempty"`
	TotalCount     *int32                      `json:"totalCount,omitempty"`
	CollectionTime *string                     `json:"collectionTime,omitempty"`
}
