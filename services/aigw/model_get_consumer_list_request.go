package aigw

type GetConsumerListRequest struct {
	InstanceId *string `json:"-"`
	PageNo     *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
	TagKey     *string `json:"-"`
	TagValue   *string `json:"-"`
	XRegion    *string `json:"-"`
}
