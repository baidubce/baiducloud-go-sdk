package aigw

type GetServiceListRequest struct {
	InstanceId    *string `json:"-"`
	ServiceSource *string `json:"-"`
	XRegion       *string `json:"-"`
}
