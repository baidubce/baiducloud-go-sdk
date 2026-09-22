package scs

type ApplicationParameterTemplateRequest struct {
	TemplateShowId   *string                   `json:"-"`
	Extra            *string                   `json:"extra,omitempty"`
	CacheClusterList []*CacheClusterShowIdItem `json:"cacheClusterList,omitempty"`
	RebootType       *int32                    `json:"rebootType,omitempty"`
	Parameters       []*Parameters             `json:"parameters,omitempty"`
}
