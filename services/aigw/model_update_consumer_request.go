package aigw

type UpdateConsumerRequest struct {
	InstanceId         *string                     `json:"-"`
	ConsumerId         *string                     `json:"-"`
	KeyType            *string                     `json:"-"`
	XRegion            *string                     `json:"-"`
	Description        *string                     `json:"description,omitempty"`
	RouteNames         []*string                   `json:"routeNames,omitempty"`
	Tags               []*Tag                      `json:"tags,omitempty"`
	CredentialOp       *CredentialOp               `json:"credentialOp,omitempty"`
	CredentialLocation *ConsumerCredentialLocation `json:"credentialLocation,omitempty"`
	IamCredential      *IAMCredentialSpec          `json:"iamCredential,omitempty"`
}
