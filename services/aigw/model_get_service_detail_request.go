package aigw

type GetServiceDetailRequest struct {
	InstanceId  *string `json:"-"`
	ServiceName *string `json:"-"`
	XRegion     *string `json:"-"`
}
