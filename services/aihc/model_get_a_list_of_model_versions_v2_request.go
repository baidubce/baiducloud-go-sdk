package aihc

type GetAListOfModelVersionsV2Request struct {
	ModelId    *string `json:"-"`
	PageNumber *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
