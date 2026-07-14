package as

type UpdateIsManagedV2Request struct {
	GroupId           *string   `json:"-"`
	UpdateIsManaged   *string   `json:"-"`
	AddManagedNodeIds []*string `json:"addManagedNodeIds,omitempty"`
	DelManagedNodeIds []*string `json:"delManagedNodeIds,omitempty"`
}
