package as

type DetachNodeV2Request struct {
	GroupId    *string   `json:"-"`
	DetachNode *string   `json:"-"`
	Nodes      []*string `json:"nodes,omitempty"`
}
