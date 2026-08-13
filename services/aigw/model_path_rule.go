package aigw

type PathRule struct {
	MatchType     *string `json:"matchType,omitempty"`
	Value         *string `json:"value,omitempty"`
	CaseSensitive *bool   `json:"caseSensitive,omitempty"`
}
