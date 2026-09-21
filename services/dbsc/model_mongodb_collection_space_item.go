package dbsc

type MongodbCollectionSpaceItem struct {
	Database                *string  `json:"database,omitempty"`
	Collection              *string  `json:"collection,omitempty"`
	StorageSize             *int64   `json:"storageSize,omitempty"`
	DataSize                *int64   `json:"dataSize,omitempty"`
	IndexSize               *int64   `json:"indexSize,omitempty"`
	ObjectCount             *int64   `json:"objectCount,omitempty"`
	IndexCount              *int32   `json:"indexCount,omitempty"`
	AvgObjSize              *int64   `json:"avgObjSize,omitempty"`
	Views                   *int32   `json:"views,omitempty"`
	FragmentationRatio      *float64 `json:"fragmentationRatio,omitempty"`
	IndexFragmentationRatio *float64 `json:"indexFragmentationRatio,omitempty"`
	TotalSize               *int64   `json:"totalSize,omitempty"`
	FreeStorageSize         *int64   `json:"freeStorageSize,omitempty"`
	IndexFreeStorageSize    *int64   `json:"indexFreeStorageSize,omitempty"`
	TotalFreeStorageSize    *int64   `json:"totalFreeStorageSize,omitempty"`
}
