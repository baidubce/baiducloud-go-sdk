package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryAssociationRelationResponse struct {
	bce.BaseResponse
	Associations []*CsnRtAssociation `json:"associations,omitempty"`
}
