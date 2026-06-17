package bcc

type EhcCluster struct {
	EhcClusterId        *string   `json:"ehcClusterId,omitempty"`
	Name                *string   `json:"name,omitempty"`
	Description         *string   `json:"description,omitempty"`
	ZoneName            *string   `json:"zoneName,omitempty"`
	CreatedTime         *string   `json:"createdTime,omitempty"`
	InstanceIds         []*string `json:"instanceIds,omitempty"`
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
}
