package as

type AttachNodeV2Request struct {
	GroupId    *string   `json:"-"`
	AttachNode *string   `json:"-"`
	Nodes      []*string `json:"nodes,omitempty"`
}
