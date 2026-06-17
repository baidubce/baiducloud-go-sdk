package bcc

type ModifyEhcClusterRequest struct {
	EhcClusterId *string `json:"ehcClusterId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
}
