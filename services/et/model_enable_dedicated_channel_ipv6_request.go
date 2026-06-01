package et

type EnableDedicatedChannelIpv6Request struct {
	EtId                *string   `json:"-"`
	EtChannelId         *string   `json:"-"`
	ClientToken         *string   `json:"-"`
	BaiduIpv6Address    *string   `json:"baiduIpv6Address,omitempty"`
	CustomerIpv6Address *string   `json:"customerIpv6Address,omitempty"`
	Ipv6Networks        []*string `json:"ipv6Networks,omitempty"`
}
