package rapidfs

type AihcResourcePoolInfo struct {
	ResourcePoolId   *string   `json:"resourcePoolId,omitempty"`
	Name             *string   `json:"name,omitempty"`
	RapidfsType      *string   `json:"type,omitempty"`
	Zones            []*string `json:"zones,omitempty"`
	BoundInstanceIds []*string `json:"boundInstanceIds,omitempty"`
}
