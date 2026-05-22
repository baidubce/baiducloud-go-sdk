package blb

type UpdateAppBlbSslListenerRequest struct {
	BlbId               *string   `json:"-"`
	ClientToken         *string   `json:"-"`
	ListenerPort        *int32    `json:"-"`
	Scheduler           *string   `json:"scheduler,omitempty"`
	CertIds             []*string `json:"certIds,omitempty"`
	EncryptionType      *string   `json:"encryptionType,omitempty"`
	EncryptionProtocols []*string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers      *string   `json:"appliedCiphers,omitempty"`
	DualAuth            *bool     `json:"dualAuth,omitempty"`
	ClientCertIds       []*string `json:"clientCertIds,omitempty"`
	Description         *string   `json:"description,omitempty"`
}
