package aigw

type VpcEndpoint struct {
	VpcEndpointId *string `json:"vpcEndpointId,omitempty"`
	VpcId         *string `json:"vpcId,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	BackendIp     *string `json:"backendIp,omitempty"`
	BackendPort   *string `json:"backendPort,omitempty"`
	EndpointIp    *string `json:"endpointIp,omitempty"`
	EndpointPort  *string `json:"endpointPort,omitempty"`
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	AigwType      *string `json:"type,omitempty"`
	Status        *string `json:"status,omitempty"`
}
