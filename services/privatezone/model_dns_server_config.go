package privatezone

type DnsServerConfig struct {
	Ip   *string `json:"ip,omitempty"`
	Port *int32  `json:"port,omitempty"`
}
