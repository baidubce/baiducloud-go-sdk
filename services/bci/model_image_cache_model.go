package bci

type ImageCacheModel struct {
	ImageCacheId       *string   `json:"imageCacheId,omitempty"`
	OriginImages       []*string `json:"originImages,omitempty"`
	Status             *string   `json:"status,omitempty"`
	Progress           *int32    `json:"progress,omitempty"`
	ExpiredTime        *string   `json:"expiredTime,omitempty"`
	CreatedTime        *string   `json:"createdTime,omitempty"`
	LastestMatchedTime *string   `json:"lastestMatchedTime,omitempty"`
}
