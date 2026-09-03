package aigw

type DeleteServiceRequest struct {
	InstanceId  *string `json:"-"`
	ServiceName *string `json:"-"`
	Namespace   *string `json:"-"`
	XRegion     *string `json:"-"`
}
