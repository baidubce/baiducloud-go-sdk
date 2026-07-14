package as

type ScalingDownV2Request struct {
	GroupId     *string   `json:"-"`
	ScalingDown *string   `json:"-"`
	Nodes       []*string `json:"nodes,omitempty"`
}
