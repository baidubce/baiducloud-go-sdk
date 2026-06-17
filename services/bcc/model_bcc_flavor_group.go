package bcc

type BccFlavorGroup struct {
	GroupId *string      `json:"groupId,omitempty"`
	Flavors []*BccFlavor `json:"flavors,omitempty"`
}
