package bcm

type Mention struct {
	BcmType *string   `json:"type,omitempty"`
	UserIds []*string `json:"userIds,omitempty"`
}
