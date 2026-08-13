package aigw

type MatcherRule struct {
	Key       *string `json:"key,omitempty"`
	MatchType *string `json:"matchType,omitempty"`
	Value     *string `json:"value,omitempty"`
}
