package bcc

type DescribeAuthorizeRulesRequest struct {
	Action    *string   `json:"-"`
	Marker    *string   `json:"marker,omitempty"`
	MaxKeys   *int32    `json:"maxKeys,omitempty"`
	RuleIds   []*string `json:"ruleIds,omitempty"`
	RuleNames []*string `json:"ruleNames,omitempty"`
}
