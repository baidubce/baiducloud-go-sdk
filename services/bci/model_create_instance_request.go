package bci

type CreateInstanceRequest struct {
	ClientToken                   *string                    `json:"-"`
	Name                          *string                    `json:"name,omitempty"`
	ZoneName                      *string                    `json:"zoneName,omitempty"`
	SecurityGroupIds              []*string                  `json:"securityGroupIds,omitempty"`
	SubnetIds                     []*string                  `json:"subnetIds,omitempty"`
	RestartPolicy                 *string                    `json:"restartPolicy,omitempty"`
	EipIp                         *string                    `json:"eipIp,omitempty"`
	AutoCreateEip                 *bool                      `json:"autoCreateEip,omitempty"`
	EipName                       *string                    `json:"eipName,omitempty"`
	EipRouteType                  *string                    `json:"eipRouteType,omitempty"`
	EipBandwidthInMbps            *int32                     `json:"eipBandwidthInMbps,omitempty"`
	EipBillingMethod              *string                    `json:"eipBillingMethod,omitempty"`
	GpuType                       *string                    `json:"gpuType,omitempty"`
	TerminationGracePeriodSeconds *int64                     `json:"terminationGracePeriodSeconds,omitempty"`
	HostName                      *string                    `json:"hostName,omitempty"`
	Tags                          []*Tag                     `json:"tags,omitempty"`
	ImageRegistryCredentials      []*ImageRegistryCredential `json:"imageRegistryCredentials,omitempty"`
	Containers                    []*Container               `json:"containers,omitempty"`
	InitContainers                []*Container               `json:"initContainers,omitempty"`
	Volume                        *Volume                    `json:"volume,omitempty"`
}
