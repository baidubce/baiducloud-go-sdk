package aigw

type IAMCredentialSpec struct {
	Name             *string   `json:"name,omitempty"`
	IamApiKeyId      *string   `json:"iamApiKeyId,omitempty"`
	IamTokenIdMasked *string   `json:"iamTokenIdMasked,omitempty"`
	IamUserId        *string   `json:"iamUserId,omitempty"`
	IamDomainId      *string   `json:"iamDomainId,omitempty"`
	ResourceIds      []*string `json:"resourceIds,omitempty"`
	InHeader         *bool     `json:"inHeader,omitempty"`
	InQuery          *bool     `json:"inQuery,omitempty"`
	KeyNames         []*string `json:"keyNames,omitempty"`
	Status           *string   `json:"status,omitempty"`
}
