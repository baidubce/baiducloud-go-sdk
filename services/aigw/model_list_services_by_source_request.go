package aigw

type ListServicesBySourceRequest struct {
	InstanceId    *string `json:"-"`
	ServiceSource *string `json:"-"`
	XRegion       *string `json:"-"`
}
