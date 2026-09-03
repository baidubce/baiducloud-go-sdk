package image

type AnimalResult struct {
	Name      *string    `json:"name,omitempty"`
	Score     *string    `json:"score,omitempty"`
	BaikeInfo *BaikeInfo `json:"baike_info,omitempty"`
}
