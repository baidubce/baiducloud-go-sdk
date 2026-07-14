package as

type BccNameConfig struct {
	BccName            *string `json:"bccName,omitempty"`
	BccHostname        *string `json:"bccHostname,omitempty"`
	AutoSeqSuffix      *bool   `json:"autoSeqSuffix,omitempty"`
	OpenHostnameDomain *bool   `json:"openHostnameDomain,omitempty"`
}
