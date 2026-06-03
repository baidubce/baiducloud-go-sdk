package ccr

type ListAcceleratorFiltersRequest struct {
	InstanceId *string `json:"-"`
	PolicyName *string `json:"-"`
	PageNo     *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
