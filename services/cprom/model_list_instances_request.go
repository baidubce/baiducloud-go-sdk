package cprom

type ListInstancesRequest struct {
	PageSize *int32  `json:"-"`
	PageNo   *int32  `json:"-"`
	OrderBy  *string `json:"-"`
	Order    *string `json:"-"`
	Phase    *string `json:"-"`
}
