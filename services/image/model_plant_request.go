package image

type PlantRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	BaikeNum *int32  `json:"baike_num,omitempty"`
}
