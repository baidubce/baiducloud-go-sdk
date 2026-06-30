package bci

type CreateImageCacheRequest struct {
	ImageCacheName       *string                    `json:"imageCacheName,omitempty"`
	OriginImages         []*OriginImage             `json:"originImages,omitempty"`
	SubnetId             *string                    `json:"subnetId,omitempty"`
	SecurityGroupId      *string                    `json:"securityGroupId,omitempty"`
	ZoneName             *string                    `json:"zoneName,omitempty"`
	TemporaryStorageSize *int32                     `json:"temporaryStorageSize,omitempty"`
	NeedEip              *bool                      `json:"needEip,omitempty"`
	EipIp                *string                    `json:"eipIp,omitempty"`
	AutoMatchImageCache  *bool                      `json:"autoMatchImageCache,omitempty"`
	ImageRegistrySecrets []*ImageRegistryCredential `json:"imageRegistrySecrets,omitempty"`
}
