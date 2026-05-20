package dns

type PublicZone struct {
	Id             *string     `json:"id,omitempty"`
	Name           *string     `json:"name,omitempty"`
	Status         *string     `json:"status,omitempty"`
	ProductVersion *string     `json:"productVersion,omitempty"`
	CreateTime     *string     `json:"createTime,omitempty"`
	ExpireTime     *string     `json:"expireTime,omitempty"`
	Tags           []*TagModel `json:"tags,omitempty"`
}
