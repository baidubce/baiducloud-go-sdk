package aigw

type CorsPolicy struct {
	Enabled          *bool          `json:"enabled,omitempty"`
	AllowOrigins     []*OriginMatch `json:"allowOrigins,omitempty"`
	AllowMethods     []*string      `json:"allowMethods,omitempty"`
	AllowHeaders     []*string      `json:"allowHeaders,omitempty"`
	ExposeHeaders    []*string      `json:"exposeHeaders,omitempty"`
	MaxAge           *int32         `json:"maxAge,omitempty"`
	AllowCredentials *bool          `json:"allowCredentials,omitempty"`
}
