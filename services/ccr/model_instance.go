package ccr

type Instance struct {
	Id            *string       `json:"id,omitempty"`
	Name          *string       `json:"name,omitempty"`
	InstanceType  *string       `json:"instanceType,omitempty"`
	PublicURL     *string       `json:"publicURL,omitempty"`
	PrivateURL    *string       `json:"privateURL,omitempty"`
	CustomDomains []*string     `json:"customDomains,omitempty"`
	Region        *string       `json:"region,omitempty"`
	Status        *string       `json:"status,omitempty"`
	CreateTime    *string       `json:"createTime,omitempty"`
	ExpireTime    *string       `json:"expireTime,omitempty"`
	Tags          []*LogicalTag `json:"tags,omitempty"`
}
