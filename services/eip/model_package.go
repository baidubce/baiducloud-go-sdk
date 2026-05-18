package eip

type Package struct {
	Id           *string `json:"id,omitempty"`
	DeductPolicy *string `json:"deductPolicy,omitempty"`
	PackageType  *string `json:"packageType,omitempty"`
	Status       *string `json:"status,omitempty"`
	Capacity     *string `json:"capacity,omitempty"`
	UsedCapacity *string `json:"usedCapacity,omitempty"`
	CreateTime   *string `json:"createTime,omitempty"`
	ActiveTime   *string `json:"activeTime,omitempty"`
	ExpireTime   *string `json:"expireTime,omitempty"`
}
