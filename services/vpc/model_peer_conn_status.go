package vpc

type PeerConnStatus struct {
	Creating      *string `json:"creating,omitempty"`
	Consulting    *string `json:"consulting,omitempty"`
	ConsultFailed *string `json:"consult_failed,omitempty"`
	Active        *string `json:"active,omitempty"`
	Down          *string `json:"down,omitempty"`
	Starting      *string `json:"starting,omitempty"`
	Stopping      *string `json:"stopping,omitempty"`
	Deleting      *string `json:"deleting,omitempty"`
	Deleted       *string `json:"deleted,omitempty"`
	Expired       *string `json:"expired,omitempty"`
	VpcError      *string `json:"error,omitempty"`
	Updating      *string `json:"updating,omitempty"`
}
