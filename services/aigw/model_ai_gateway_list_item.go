package aigw

type AIGatewayListItem struct {
	InstanceId              *string      `json:"instanceId,omitempty"`
	Name                    *string      `json:"name,omitempty"`
	IngressStatus           *string      `json:"ingressStatus,omitempty"`
	InternalIP              *string      `json:"internalIP,omitempty"`
	ExternalIP              *string      `json:"externalIP,omitempty"`
	CreateTime              *string      `json:"createTime,omitempty"`
	Region                  *string      `json:"region,omitempty"`
	Replicas                *int32       `json:"replicas,omitempty"`
	InstallMode             *string      `json:"installMode,omitempty"`
	VpcCidr                 *string      `json:"vpcCidr,omitempty"`
	VpcId                   *string      `json:"vpcId,omitempty"`
	SubnetId                *string      `json:"subnetId,omitempty"`
	GatewayType             *string      `json:"gatewayType,omitempty"`
	PublicAccessible        *bool        `json:"publicAccessible,omitempty"`
	DeleteProtection        *bool        `json:"deleteProtection,omitempty"`
	Description             *string      `json:"description,omitempty"`
	Namespace               *string      `json:"namespace,omitempty"`
	EnableIngress           *bool        `json:"enableIngress,omitempty"`
	EnableAllIngressClasses *bool        `json:"enableAllIngressClasses,omitempty"`
	EnableAllNamespaces     *bool        `json:"enableAllNamespaces,omitempty"`
	IngressClasses          []*string    `json:"ingressClasses,omitempty"`
	WatchNamespaces         []*string    `json:"watchNamespaces,omitempty"`
	BaEndpoint              *VpcEndpoint `json:"baEndpoint,omitempty"`
	AssociatedCluster       *string      `json:"associatedCluster,omitempty"`
	SrcProduct              *string      `json:"srcProduct,omitempty"`
	BlbLongId               *string      `json:"blbLongId,omitempty"`
	WafId                   *string      `json:"wafId,omitempty"`
	WafEnable               *bool        `json:"wafEnable,omitempty"`
	PrivateDomainName       *string      `json:"privateDomainName,omitempty"`
	PublicDomainName        *string      `json:"publicDomainName,omitempty"`
	NetworkType             *string      `json:"networkType,omitempty"`
	DomainStatus            *string      `json:"domainStatus,omitempty"`
	SecurityGroupId         *string      `json:"securityGroupId,omitempty"`
	Tags                    []*Tag       `json:"tags,omitempty"`
	Version                 *string      `json:"version,omitempty"`
	AihcArgs                *AihcArgs    `json:"aihcArgs,omitempty"`
}
