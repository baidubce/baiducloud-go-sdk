package bls

type ListFastQueryRequest struct {
	Project      *string `json:"-"`
	LogStoreName *string `json:"-"`
	NamePattern  *string `json:"-"`
	Order        *string `json:"-"`
	OrderBy      *string `json:"-"`
	PageNo       *int32  `json:"-"`
	PageSize     *int32  `json:"-"`
}
