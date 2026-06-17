package bcc

type ReservedInstance struct {
	ReservedInstanceId   *string `json:"reservedInstanceId,omitempty"`
	ZoneName             *string `json:"zoneName,omitempty"`
	ReservedInstanceName *string `json:"reservedInstanceName,omitempty"`
	EhcClusterId         *string `json:"ehcClusterId,omitempty"`
}
