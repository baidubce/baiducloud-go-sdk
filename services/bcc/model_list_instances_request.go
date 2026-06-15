package bcc

type ListInstancesRequest struct {
	Marker            *string `json:"-"`
	MaxKeys           *int32  `json:"-"`
	InternalIp        *string `json:"-"`
	DedicatedHostId   *string `json:"-"`
	ZoneName          *string `json:"-"`
	ShowRdmaTopo      *string `json:"-"`
	InstanceIds       *string `json:"-"`
	InstanceNames     *string `json:"-"`
	FuzzyInstanceName *string `json:"-"`
	VolumeIds         *string `json:"-"`
	DeploySetIds      *string `json:"-"`
	SecurityGroupIds  *string `json:"-"`
	PaymentTiming     *string `json:"-"`
	Status            *string `json:"-"`
	Tags              *string `json:"-"`
	VpcId             *string `json:"-"`
	PrivateIps        *string `json:"-"`
	EhcClusterId      *string `json:"-"`
}
