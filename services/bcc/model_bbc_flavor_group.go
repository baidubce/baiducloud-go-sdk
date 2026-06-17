package bcc

type BbcFlavorGroup struct {
	GroupId *string      `json:"groupId,omitempty"`
	Flavors []*BbcFlavor `json:"flavors,omitempty"`
}
