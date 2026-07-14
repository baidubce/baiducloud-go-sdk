package as

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetAsGroupV2Response struct {
	bce.BaseResponse
	GroupId           *string            `json:"groupId,omitempty"`
	GroupName         *string            `json:"groupName,omitempty"`
	Region            *string            `json:"region,omitempty"`
	Status            *string            `json:"status,omitempty"`
	VpcInfo           *VpcInfo           `json:"vpcInfo,omitempty"`
	ZoneInfo          []*interface{}     `json:"zoneInfo,omitempty"`
	Config            *GroupConfig       `json:"config,omitempty"`
	KeypairId         *string            `json:"keypairId,omitempty"`
	KeypairName       *string            `json:"keypairName,omitempty"`
	KeepImageLogin    *bool              `json:"keepImageLogin,omitempty"`
	Blb               []*BlbInfo         `json:"blb,omitempty"`
	BlbUnbindWaitTime *int32             `json:"blbUnbindWaitTime,omitempty"`
	NodeNum           *int32             `json:"nodeNum,omitempty"`
	CreateTime        *string            `json:"createTime,omitempty"`
	RdsIds            *string            `json:"rdsIds,omitempty"`
	ScsIds            *string            `json:"scsIds,omitempty"`
	ExpansionStrategy *string            `json:"expansionStrategy,omitempty"`
	ShrinkageStrategy *string            `json:"shrinkageStrategy,omitempty"`
	RelationTag       *bool              `json:"relationTag,omitempty"`
	Tags              []*TagInfo         `json:"tags,omitempty"`
	CmdConfig         *CmdConfig         `json:"cmdConfig,omitempty"`
	BccNameConfig     *BccNameConfig     `json:"bccNameConfig,omitempty"`
	EipConfig         *EipConfig         `json:"eipConfig,omitempty"`
	HealthCheck       *HealthCheckConfig `json:"healthCheck,omitempty"`
}
