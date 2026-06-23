package cprom

type BindClusterRequest struct {
	InstanceId *string `json:"-"`
	Action     *string `json:"-"`
	ClusterId  *string `json:"clusterId,omitempty"`
}
