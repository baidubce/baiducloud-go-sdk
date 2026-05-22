package blb

type AppListenerModel struct {
	ListenerPort          *int32    `json:"listenerPort,omitempty"`
	ListenerType          *string   `json:"listenerType,omitempty"`
	Scheduler             *string   `json:"scheduler,omitempty"`
	TcpSessionTimeout     *int32    `json:"tcpSessionTimeout,omitempty"`
	UdpSessionTimeout     *int32    `json:"udpSessionTimeout,omitempty"`
	KeepSession           *bool     `json:"keepSession,omitempty"`
	KeepSessionType       *string   `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    *int32    `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName *string   `json:"keepSessionCookieName,omitempty"`
	XForwardFor           *bool     `json:"xForwardFor,omitempty"`
	XForwardedProto       *bool     `json:"xForwardedProto,omitempty"`
	ServerTimeout         *int32    `json:"serverTimeout,omitempty"`
	RedirectPort          *int32    `json:"redirectPort,omitempty"`
	CertIds               []*string `json:"certIds,omitempty"`
	DualAuth              *bool     `json:"dualAuth,omitempty"`
	ClientCertIds         []*string `json:"clientCertIds,omitempty"`
	EncryptionType        *string   `json:"encryptionType,omitempty"`
	EncryptionProtocols   []*string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers        *string   `json:"appliedCiphers,omitempty"`
}
