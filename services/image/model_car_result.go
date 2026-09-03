package image

type CarResult struct {
	Name      *string    `json:"name,omitempty"`
	Score     *float64   `json:"score,omitempty"`
	Year      *string    `json:"year,omitempty"`
	BaikeInfo *BaikeInfo `json:"baike_info,omitempty"`
}
