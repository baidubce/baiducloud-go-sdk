package ccr

type ArtifactScanOverview struct {
	Description *string   `json:"description,omitempty"`
	FixVersion  *string   `json:"fixVersion,omitempty"`
	Id          *string   `json:"id,omitempty"`
	Links       []*string `json:"links,omitempty"`
	CcrPackage  *string   `json:"package,omitempty"`
	Severity    *string   `json:"severity,omitempty"`
	Version     *string   `json:"version,omitempty"`
}
