package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListBaseDdosAttackRecordResponse struct {
	bce.BaseResponse
	AttackRecordList []*DdosAttackRecordModel `json:"attackRecordList,omitempty"`
}
