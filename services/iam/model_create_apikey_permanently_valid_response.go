package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateApikeyPermanentlyValidResponse struct {
	bce.BaseResponse
	Id         *string   `json:"id,omitempty"`
	TokenId    *string   `json:"tokenId,omitempty"`
	Name       *string   `json:"name,omitempty"`
	UserId     *string   `json:"userId,omitempty"`
	Services   []*string `json:"services,omitempty"`
	CreateTime *string   `json:"createTime,omitempty"`
	UpdateTime *string   `json:"updateTime,omitempty"`
	DomainId   *string   `json:"domainId,omitempty"`
	Acl        *ACL      `json:"acl,omitempty"`
}
