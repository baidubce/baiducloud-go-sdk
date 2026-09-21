package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbCollectionSpaceResponse struct {
	bce.BaseResponse
	Items          []*MongodbCollectionSpaceItem `json:"items,omitempty"`
	TotalCount     *int32                        `json:"totalCount,omitempty"`
	CollectionTime *string                       `json:"collectionTime,omitempty"`
}
