package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateConsumerResponse struct {
	bce.BaseResponse
	Success        *bool                     `json:"success,omitempty"`
	Status         *int32                    `json:"status,omitempty"`
	ConsumerId     *string                   `json:"consumerId,omitempty"`
	Credential     *string                   `json:"credential,omitempty"`
	Credentials    []*ConsumerCredentialInfo `json:"credentials,omitempty"`
	CredentialType *string                   `json:"credentialType,omitempty"`
}
