package as

type UpdateProtectV2Request struct {
	GroupId       *string   `json:"-"`
	UpdateProtect *string   `json:"-"`
	Nodes         []*string `json:"nodes,omitempty"`
	IsProtected   *bool     `json:"isProtected,omitempty"`
}
