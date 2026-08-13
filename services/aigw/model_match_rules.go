package aigw

type MatchRules struct {
	PathRule    *interface{} `json:"pathRule,omitempty"`
	Methods     []*string    `json:"methods,omitempty"`
	Headers     *interface{} `json:"headers,omitempty"`
	QueryParams *interface{} `json:"queryParams,omitempty"`
}
