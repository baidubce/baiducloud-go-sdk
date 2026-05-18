package eip

type AddEipGroupCountRequest struct {
	Id            *string `json:"-"`
	ClientToken   *string `json:"-"`
	EipAddCount   *int32  `json:"eipAddCount,omitempty"`
	Eipv6AddCount *int32  `json:"eipv6AddCount,omitempty"`
}
