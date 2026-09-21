package dbsc

type GetMongodbCollectionIndexesRequest struct {
	Product    *string `json:"-"`
	AppId      *string `json:"-"`
	NodeId     *string `json:"-"`
	Database   *string `json:"-"`
	Collection *string `json:"-"`
}
