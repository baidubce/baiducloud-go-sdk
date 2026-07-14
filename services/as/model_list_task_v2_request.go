package as

type ListTaskV2Request struct {
	Groupid   *string `json:"-"`
	Order     *string `json:"-"`
	OrderBy   *string `json:"-"`
	PageNo    *int32  `json:"-"`
	PageSize  *int32  `json:"-"`
	StartTime *string `json:"-"`
	EndTime   *string `json:"-"`
}
