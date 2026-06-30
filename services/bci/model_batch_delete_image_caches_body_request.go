package bci

type BatchDeleteImageCachesRequest struct {
	ImageCacheIds []*string `json:"imageCacheIds,omitempty"`
}
