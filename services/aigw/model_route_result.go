package aigw

type RouteResult struct {
	RouteName                   *string           `json:"routeName,omitempty"`
	SrcProduct                  *string           `json:"srcProduct,omitempty"`
	AccessMode                  *string           `json:"accessMode,omitempty"`
	Domains                     []*string         `json:"domains,omitempty"`
	WebDomains                  []*string         `json:"webDomains,omitempty"`
	WebSubdomain                *string           `json:"webSubdomain,omitempty"`
	ServicePath                 *string           `json:"servicePath,omitempty"`
	CreateTime                  *string           `json:"createTime,omitempty"`
	UpdateTime                  *string           `json:"updateTime,omitempty"`
	MatchRules                  *MatchRules       `json:"matchRules,omitempty"`
	MultiService                *bool             `json:"multiService,omitempty"`
	TrafficDistributionStrategy *string           `json:"trafficDistributionStrategy,omitempty"`
	EnableWeightAdjust          *bool             `json:"enableWeightAdjust,omitempty"`
	TargetService               *TargetService    `json:"targetService,omitempty"`
	Rewrite                     *Rewrite          `json:"rewrite,omitempty"`
	RegexRewrite                *RegexRewrite     `json:"regexRewrite,omitempty"`
	CustomHeaders               []*CustomHeader   `json:"customHeaders,omitempty"`
	AuthEnabled                 *bool             `json:"authEnabled,omitempty"`
	AllowedConsumers            []*string         `json:"allowedConsumers,omitempty"`
	TokenRateLimit              *TokenRateLimit   `json:"tokenRateLimit,omitempty"`
	RequestRateLimit            *RequestRateLimit `json:"requestRateLimit,omitempty"`
	TimeoutPolicy               *TimeoutPolicy    `json:"timeoutPolicy,omitempty"`
	RetryPolicy                 *RetryPolicy      `json:"retryPolicy,omitempty"`
	CorsPolicy                  *CorsPolicy       `json:"corsPolicy,omitempty"`
	ResponseHeaders             *ResponseHeaders  `json:"responseHeaders,omitempty"`
	FallbackConfig              *FallbackConfig   `json:"fallbackConfig,omitempty"`
}
