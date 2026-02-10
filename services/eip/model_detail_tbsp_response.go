package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DetailTbspResponse struct {
	bce.BaseResponse
	Name                      *string                  `json:"name,omitempty"`
	Id                        *string                  `json:"id,omitempty"`
	DefenseLineType           *string                  `json:"defenseLineType,omitempty"`
	DefenseCountQuota         *int32                   `json:"defenseCountQuota,omitempty"`
	IpList                    []*TbspIpModel           `json:"ipList,omitempty"`
	IpTotalCount              *int32                   `json:"ipTotalCount,omitempty"`
	AutoRenewSwitch           *int32                   `json:"autoRenewSwitch,omitempty"`
	ProductStatus             *string                  `json:"productStatus,omitempty"`
	CreateTime                *string                  `json:"createTime,omitempty"`
	ExpireTime                *string                  `json:"expireTime,omitempty"`
	DefenseEnable             *int32                   `json:"defenseEnable,omitempty"`
	AttackingRecordList       []*TbspAttackRecordModel `json:"attackingRecordList,omitempty"`
	AttackingRecordTotalCount *int32                   `json:"attackingRecordTotalCount,omitempty"`
	Tags                      []*TagModel              `json:"tags,omitempty"`
}
