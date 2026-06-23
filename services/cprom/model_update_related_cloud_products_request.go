package cprom

type UpdateRelatedCloudProductsRequest struct {
	InstanceId *string   `json:"-"`
	Scopes     []*string `json:"scopes,omitempty"`
}
