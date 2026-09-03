package aigw

type CreateConsumerRequest struct {
	InstanceId     *string                 `json:"-"`
	XRegion        *string                 `json:"-"`
	ConsumerName   *string                 `json:"consumerName,omitempty"`
	Description    *string                 `json:"description,omitempty"`
	AuthType       *string                 `json:"authType,omitempty"`
	CredentialType *string                 `json:"credentialType,omitempty"`
	RouteNames     []*string               `json:"routeNames,omitempty"`
	Tags           []*Tag                  `json:"tags,omitempty"`
	Credential     *ConsumerCredentialSpec `json:"credential,omitempty"`
	IamCredential  *IAMCredentialSpec      `json:"iamCredential,omitempty"`
}
