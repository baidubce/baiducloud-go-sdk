package ccr

type InstanceQuota struct {
	Repo      *int32 `json:"repo,omitempty"`
	Chart     *int32 `json:"chart,omitempty"`
	Namespace *int32 `json:"namespace,omitempty"`
	Vpc       *int32 `json:"vpc,omitempty"`
}
