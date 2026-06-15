package bcm

type Receiver struct {
	ReceiversId       *string `json:"receivers.id,omitempty"`
	ReceiversDomainId *string `json:"receivers.domainId,omitempty"`
	ReceiversName     *string `json:"receivers.name,omitempty"`
	ReceiversEmail    *string `json:"receivers.email,omitempty"`
	ReceiversPhone    *string `json:"receivers.phone,omitempty"`
}
