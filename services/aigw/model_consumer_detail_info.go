package aigw

type ConsumerDetailInfo struct {
	ConsumerId     *string                   `json:"consumerId,omitempty"`
	ConsumerName   *string                   `json:"consumerName,omitempty"`
	Description    *string                   `json:"description,omitempty"`
	AuthType       *string                   `json:"authType,omitempty"`
	CredentialType *string                   `json:"credentialType,omitempty"`
	RouteNames     []*string                 `json:"routeNames,omitempty"`
	Tags           []*Tag                    `json:"tags,omitempty"`
	Credentials    []*ConsumerCredentialInfo `json:"credentials,omitempty"`
	IamCredential  *IAMCredentialSpec        `json:"iamCredential,omitempty"`
}
