package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetServiceDetailResponse struct {
	bce.BaseResponse
	ClusterId           *string   `json:"clusterId,omitempty"`
	ClusterIds          []*string `json:"clusterIds,omitempty"`
	Namespace           *string   `json:"namespace,omitempty"`
	RouteCount          *int32    `json:"routeCount,omitempty"`
	ServiceSource       *string   `json:"serviceSource,omitempty"`
	ServiceStatus       *string   `json:"serviceStatus,omitempty"`
	ServicePort         []*string `json:"servicePort,omitempty"`
	ServiceAddresses    []*string `json:"serviceAddresses,omitempty"`
	ServiceProtocol     *string   `json:"serviceProtocol,omitempty"`
	McpServerHosts      []*string `json:"mcpServerHosts,omitempty"`
	Provider            *string   `json:"provider,omitempty"`
	Endpoint            *string   `json:"endpoint,omitempty"`
	ApiKeys             []*string `json:"apiKeys,omitempty"`
	FailoverEnabled     *bool     `json:"failoverEnabled,omitempty"`
	FailoverModel       *string   `json:"failoverModel,omitempty"`
	FailureThreshold    *int64    `json:"failureThreshold,omitempty"`
	HealthCheckInterval *int64    `json:"healthCheckInterval,omitempty"`
	HealthCheckTimeout  *int64    `json:"healthCheckTimeout,omitempty"`
	CredentialSource    *string   `json:"credentialSource,omitempty"`
	CredentialNames     []*string `json:"credentialNames,omitempty"`
}
