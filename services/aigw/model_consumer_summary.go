package aigw

type ConsumerSummary struct {
	ConsumerId     *string   `json:"consumerId,omitempty"`
	ConsumerName   *string   `json:"consumerName,omitempty"`
	Description    *string   `json:"description,omitempty"`
	AuthType       *string   `json:"authType,omitempty"`
	CredentialType *string   `json:"credentialType,omitempty"`
	Tags           []*Tag    `json:"tags,omitempty"`
	RouteNames     []*string `json:"routeNames,omitempty"`
}
