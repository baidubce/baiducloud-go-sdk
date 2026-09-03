package image

type DishRequest struct {
	Image           *string  `json:"image,omitempty"`
	Url             *string  `json:"url,omitempty"`
	TopNum          *int32   `json:"top_num,omitempty"`
	FilterThreshold *float32 `json:"filter_threshold,omitempty"`
	BaikeNum        *int32   `json:"baike_num,omitempty"`
}
