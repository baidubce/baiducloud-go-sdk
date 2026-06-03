package ccr

type HelmChartVersion struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	ApiVersion  *string   `json:"apiVersion,omitempty"`
	AppVersion  *string   `json:"appVersion,omitempty"`
	Version     *string   `json:"version,omitempty"`
	Urls        []*string `json:"urls,omitempty"`
	Digest      *string   `json:"digest,omitempty"`
	Engine      *string   `json:"engine,omitempty"`
	Home        *string   `json:"home,omitempty"`
	Icon        *string   `json:"icon,omitempty"`
	Sources     []*string `json:"sources,omitempty"`
	Created     *string   `json:"created,omitempty"`
	Deprecated  *bool     `json:"deprecated,omitempty"`
	Removed     *bool     `json:"removed,omitempty"`
	Maintainers []*string `json:"maintainers,omitempty"`
}
