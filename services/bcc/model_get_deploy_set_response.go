package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDeploySetResponse struct {
	bce.BaseResponse
	ShortId               *string                    `json:"shortId,omitempty"`
	Uuid                  *string                    `json:"uuid,omitempty"`
	InstanceTotal         *int32                     `json:"instanceTotal,omitempty"`
	InstanceCount         *int32                     `json:"instanceCount,omitempty"`
	BccInstanceCnt        *int32                     `json:"bccInstanceCnt,omitempty"`
	BbcInstanceCnt        *int32                     `json:"bbcInstanceCnt,omitempty"`
	AzIntstanceStatisList []*AzIntstanceStatisDetail `json:"azIntstanceStatisList,omitempty"`
}
