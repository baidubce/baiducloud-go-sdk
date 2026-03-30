package vpc

type GlrItem struct {
	GlrId          *string `json:"glrId,omitempty"`
	IpVersion      *string `json:"ipVersion,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	ServiceType    *string `json:"serviceType,omitempty"`
	SubServiceType *string `json:"subServiceType,omitempty"`
	ResourceId     *string `json:"resourceId,omitempty"`
	Direction      *string `json:"direction,omitempty"`
	Cidr           *string `json:"cidr,omitempty"`
	Bandwidth      *int32  `json:"bandwidth,omitempty"`
	Enable         *bool   `json:"enable,omitempty"`
	CreatedTime    *string `json:"createdTime,omitempty"`
	UpdatedTime    *string `json:"updatedTime,omitempty"`
}
