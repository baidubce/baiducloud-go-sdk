package scs

type ProxyNodeReplaceRequest struct {
	InstanceId *string   `json:"-"`
	ProxyList  []*string `json:"proxyList,omitempty"`
	ScsDefer   *bool     `json:"defer,omitempty"`
}
