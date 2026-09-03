package image

type CarRequest struct {
	Image       *string `json:"image,omitempty"`
	Url         *string `json:"url,omitempty"`
	TopNum      *int32  `json:"top_num,omitempty"`
	BaikeNum    *int32  `json:"baike_num,omitempty"`
	OutputBrand *bool   `json:"output_brand,omitempty"`
}
