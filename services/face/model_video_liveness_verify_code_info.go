package face

type VideoLivenessVerifyCodeInfo struct {
	Create     *string  `json:"create,omitempty"`
	Identify   *string  `json:"identify,omitempty"`
	Similarity *float32 `json:"similarity,omitempty"`
}
