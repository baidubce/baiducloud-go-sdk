package ccr

type HelmChart struct {
	Name          *string `json:"name,omitempty"`
	TotalVersions *int32  `json:"totalVersions,omitempty"`
	LatestVersion *string `json:"latestVersion,omitempty"`
	Home          *string `json:"home,omitempty"`
	Icon          *string `json:"icon,omitempty"`
	Deprecated    *bool   `json:"deprecated,omitempty"`
	Created       *string `json:"created,omitempty"`
	Updated       *string `json:"updated,omitempty"`
}
