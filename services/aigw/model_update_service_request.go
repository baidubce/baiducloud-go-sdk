package aigw

type UpdateServiceRequest struct {
	InstanceId       *string   `json:"-"`
	ServiceNamePath  *string   `json:"-"`
	XRegion          *string   `json:"-"`
	ServiceName      *string   `json:"serviceName,omitempty"`
	ServiceAddresses []*string `json:"serviceAddresses,omitempty"`
	ServiceProtocol  *string   `json:"serviceProtocol,omitempty"`
	Provider         *string   `json:"provider,omitempty"`
	Endpoint         *string   `json:"endpoint,omitempty"`
	ApiKeys          []*string `json:"apiKeys,omitempty"`
	FailoverEnabled  *bool     `json:"failoverEnabled,omitempty"`
	FailoverModel    *string   `json:"failoverModel,omitempty"`
	CredentialSource *string   `json:"credentialSource,omitempty"`
	CredentialNames  []*string `json:"credentialNames,omitempty"`
}
