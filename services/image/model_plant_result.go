package image

type PlantResult struct {
	Name      *string    `json:"name,omitempty"`
	Score     *float32   `json:"score,omitempty"`
	BaikeInfo *BaikeInfo `json:"baike_info,omitempty"`
}
