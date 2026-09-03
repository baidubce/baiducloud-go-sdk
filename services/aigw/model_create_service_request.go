package aigw

type CreateServiceRequest struct {
	InstanceId       *string        `json:"-"`
	XRegion          *string        `json:"-"`
	ServiceSource    *string        `json:"serviceSource,omitempty"`
	Namespace        *string        `json:"namespace,omitempty"`
	ServiceName      *string        `json:"serviceName,omitempty"`
	ClusterId        *string        `json:"clusterId,omitempty"`
	ClusterIds       []*string      `json:"clusterIds,omitempty"`
	ServiceList      []*ServiceItem `json:"serviceList,omitempty"`
	RegistryId       *string        `json:"registryId,omitempty"`
	ServiceAddresses []*string      `json:"serviceAddresses,omitempty"`
	ServiceProtocol  *string        `json:"serviceProtocol,omitempty"`
	Provider         *string        `json:"provider,omitempty"`
	Endpoint         *string        `json:"endpoint,omitempty"`
	ApiKeys          []*string      `json:"apiKeys,omitempty"`
	CredentialSource *string        `json:"credentialSource,omitempty"`
	CredentialNames  []*string      `json:"credentialNames,omitempty"`
	FailoverEnabled  *bool          `json:"failoverEnabled,omitempty"`
	FailoverModel    *string        `json:"failoverModel,omitempty"`
}
