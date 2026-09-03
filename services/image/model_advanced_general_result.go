package image

type AdvancedGeneralResult struct {
	Keyword   *string    `json:"keyword,omitempty"`
	Score     *float32   `json:"score,omitempty"`
	Root      *string    `json:"root,omitempty"`
	BaikeInfo *BaikeInfo `json:"baike_info,omitempty"`
}
