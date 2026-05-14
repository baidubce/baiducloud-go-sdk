package snic

type Endpoint struct {
	EndpointId  *string `json:"endpointId,omitempty"`
	Name        *string `json:"name,omitempty"`
	IpAddress   *string `json:"ipAddress,omitempty"`
	Status      *string `json:"status,omitempty"`
	Service     *string `json:"service,omitempty"`
	SubnetId    *string `json:"subnetId,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	VpcId       *string `json:"vpcId,omitempty"`
	ProductType *string `json:"productType,omitempty"`
}
