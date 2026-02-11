

package eip




type EipModel struct {
	Name *string `json:"name,omitempty"`
	Eip *string `json:"eip,omitempty"`
	EipId *string `json:"eipId,omitempty"`
	Status *string `json:"status,omitempty"`
	EipInstanceType *string `json:"eipInstanceType,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	ShareGroupId *string `json:"shareGroupId,omitempty"`
	DefaultDomesticBandwidth *int32 `json:"defaultDomesticBandwidth,omitempty"`
	BandwidthInMbps *int32 `json:"bandwidthInMbps,omitempty"`
	BwShortId *string `json:"bwShortId,omitempty"`
	BwBandwidthInMbps *int32 `json:"bwBandwidthInMbps,omitempty"`
	DomesticBwShortId *string `json:"domesticBwShortId,omitempty"`
	DomesticBwBandwidthInMbps *int32 `json:"domesticBwBandwidthInMbps,omitempty"`
	PaymentTiming *string `json:"paymentTiming,omitempty"`
	BillingMethod *string `json:"billingMethod,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty"`
	Region *string `json:"region,omitempty"`
	RouteType *string `json:"routeType,omitempty"`
	Tags []*TagModel `json:"tags,omitempty"`
	DeleteProtect *bool `json:"deleteProtect,omitempty"`
	NativeGroup *bool `json:"nativeGroup,omitempty"`
	OriginalBandwidth *int32 `json:"originalBandwidth,omitempty"`
	OriginProductType *string `json:"originProductType,omitempty"`
	OriginSubProductType *string `json:"originSubProductType,omitempty"`
}

