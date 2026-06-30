package bci

type ListImageCachesRequest struct {
	PageSize *int32 `json:"-"`
	PageNo   *int32 `json:"-"`
}
