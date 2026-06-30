package bci

type ListInstancesRequest struct {
	KeywordType *string `json:"-"`
	Keyword     *string `json:"-"`
	Marker      *string `json:"-"`
	MaxKeys     *int32  `json:"-"`
}
