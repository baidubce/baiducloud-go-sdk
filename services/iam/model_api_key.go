package iam

type APIKey struct {
	Id         *string `json:"id,omitempty"`
	TokenId    *string `json:"tokenId,omitempty"`
	Name       *string `json:"name,omitempty"`
	UserId     *string `json:"userId,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
}
