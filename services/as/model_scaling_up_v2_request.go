package as

type ScalingUpV2Request struct {
	GroupId           *string   `json:"-"`
	ScalingUp         *string   `json:"-"`
	NodeCount         *int32    `json:"nodeCount,omitempty"`
	Zone              []*string `json:"zone,omitempty"`
	ExpansionStrategy *string   `json:"expansionStrategy,omitempty"`
}
