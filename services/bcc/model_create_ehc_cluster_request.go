package bcc

type CreateEhcClusterRequest struct {
	Name        *string `json:"name,omitempty"`
	ZoneName    *string `json:"zoneName,omitempty"`
	Description *string `json:"description,omitempty"`
}
