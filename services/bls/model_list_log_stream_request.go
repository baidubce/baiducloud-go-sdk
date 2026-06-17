package bls

type ListLogStreamRequest struct {
	LogStoreName *string `json:"-"`
	Project      *string `json:"-"`
	NamePattern  *string `json:"-"`
	Order        *string `json:"-"`
	OrderBy      *string `json:"-"`
	PageNo       *int32  `json:"-"`
	PageSize     *int32  `json:"-"`
}
