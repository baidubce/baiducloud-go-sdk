package dns

type UpgradeTheFreeDomainNameToTheUniversalVersionRequest struct {
	Action      *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Names       []*string `json:"names,omitempty"`
	Billing     *Billing  `json:"billing,omitempty"`
}
