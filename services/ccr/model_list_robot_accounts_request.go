package ccr

type ListRobotAccountsRequest struct {
	InstanceId *string `json:"-"`
	Status     *string `json:"-"`
	PageNo     *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
