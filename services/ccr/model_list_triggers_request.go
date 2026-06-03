package ccr

type ListTriggersRequest struct {
	InstanceId *string `json:"-"`
	PolicyName *string `json:"-"`
	PageNo     *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
