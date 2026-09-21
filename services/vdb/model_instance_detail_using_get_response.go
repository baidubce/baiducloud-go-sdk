package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type InstanceDetailUsingGETResponse struct {
	bce.BaseResponse
	AutoRenewRule         *VdbAutoRenewRule  `json:"autoRenewRule,omitempty"`
	AvailabilityZone      *string            `json:"availabilityZone,omitempty"`
	AzInfos               []*AzInfo          `json:"azInfos,omitempty"`
	BcmCycle              *int32             `json:"bcmCycle,omitempty"`
	Components            []*MilvusComponent `json:"components,omitempty"`
	CreateTime            *string            `json:"createTime,omitempty"`
	DataNodeNum           *int32             `json:"dataNodeNum,omitempty"`
	DataNodes             []*DataNode        `json:"dataNodes,omitempty"`
	DataStatus            *string            `json:"dataStatus,omitempty"`
	DiskType              *string            `json:"diskType,omitempty"`
	Domain                *string            `json:"domain,omitempty"`
	Eip                   *string            `json:"eip,omitempty"`
	EnableEmbedding       *bool              `json:"enableEmbedding,omitempty"`
	EnableEncryption      *bool              `json:"enableEncryption,omitempty"`
	EnableTDE             *string            `json:"enableTDE,omitempty"`
	EngineMinorVersion    *string            `json:"engineMinorVersion,omitempty"`
	EngineType            *string            `json:"engineType,omitempty"`
	EngineVersion         *string            `json:"engineVersion,omitempty"`
	ExpireDate            *int32             `json:"expireDate,omitempty"`
	InstanceExpireTime    *string            `json:"instanceExpireTime,omitempty"`
	InstanceId            *string            `json:"instanceId,omitempty"`
	InstanceName          *string            `json:"instanceName,omitempty"`
	InstanceType          *string            `json:"instanceType,omitempty"`
	Ip                    *string            `json:"ip,omitempty"`
	LogService            *LogService        `json:"logService,omitempty"`
	LoggingServices       []*LoggingService  `json:"loggingServices,omitempty"`
	NodeSpec              *string            `json:"nodeSpec,omitempty"`
	Nodes                 []*MilvusNode      `json:"nodes,omitempty"`
	OrderStatus           *string            `json:"orderStatus,omitempty"`
	PackageVersion        *string            `json:"packageVersion,omitempty"`
	Port                  *int32             `json:"port,omitempty"`
	ProductType           *string            `json:"productType,omitempty"`
	Proxies               []*Proxy           `json:"proxies,omitempty"`
	ProxyNodeSpec         *string            `json:"proxyNodeSpec,omitempty"`
	ProxyNum              *int32             `json:"proxyNum,omitempty"`
	Status                *string            `json:"status,omitempty"`
	Subnets               []*Subnet          `json:"subnets,omitempty"`
	SupportEmbedding      *bool              `json:"supportEmbedding,omitempty"`
	TargetPackage         *TargetPackage     `json:"targetPackage,omitempty"`
	TotalDiskCapacityInGB *int32             `json:"totalDiskCapacityInGB,omitempty"`
	TotalMemCapacityInGB  *int32             `json:"totalMemCapacityInGB,omitempty"`
	Upgradable            *bool              `json:"upgradable,omitempty"`
	UsedDiskCapacityInGB  *float64           `json:"usedDiskCapacityInGB,omitempty"`
	UsedMemCapacityInGB   *float64           `json:"usedMemCapacityInGB,omitempty"`
	Vip                   *string            `json:"vip,omitempty"`
	VpcCidr               *string            `json:"vpcCidr,omitempty"`
	VpcId                 *string            `json:"vpcId,omitempty"`
	VpcName               *string            `json:"vpcName,omitempty"`
}
