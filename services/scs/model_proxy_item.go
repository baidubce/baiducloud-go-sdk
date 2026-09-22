package scs

type ProxyItem struct {
	Uuid             *string `json:"uuid,omitempty"`
	NodeShowId       *string `json:"nodeShowId,omitempty"`
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	NodeId           *string `json:"nodeId,omitempty"`
}
