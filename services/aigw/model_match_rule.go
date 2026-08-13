package aigw

type MatchRule struct {
	PathRule    *PathRule      `json:"pathRule,omitempty"`
	Methods     []*string      `json:"methods,omitempty"`
	Headers     []*MatcherRule `json:"headers,omitempty"`
	QueryParams []*MatcherRule `json:"queryParams,omitempty"`
}
