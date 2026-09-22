package face

type VideoLivenessVerifyResult struct {
	Score         *float32                       `json:"score,omitempty"`
	Maxspoofing   *float32                       `json:"maxspoofing,omitempty"`
	SpoofingScore *float32                       `json:"spoofing_score,omitempty"`
	Thresholds    *VideoLivenessVerifyThresholds `json:"thresholds,omitempty"`
	Code          *VideoLivenessVerifyCodeInfo   `json:"code,omitempty"`
	LipLanguage   *string                        `json:"lip_language,omitempty"`
	ActionVerify  *string                        `json:"action_verify,omitempty"`
	BestImage     *BestImage                     `json:"best_image,omitempty"`
	PicList       []*PicItem                     `json:"pic_list,omitempty"`
}
