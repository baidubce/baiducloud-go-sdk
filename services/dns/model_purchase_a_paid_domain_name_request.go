package dns

type PurchaseAPaidDomainNameRequest struct {
	ClientToken    *string   `json:"-"`
	Names          []*string `json:"names,omitempty"`
	ProductVersion *string   `json:"productVersion,omitempty"`
	Billing        *Billing  `json:"billing,omitempty"`
}
