package cfw

type StatelessCfw struct {
	CfwId           *string   `json:"cfwId,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Description     *string   `json:"description,omitempty"`
	CreatedTime     *string   `json:"createdTime,omitempty"`
	BindInstanceNum *int32    `json:"bindInstanceNum,omitempty"`
	CfwType         *int32    `json:"type,omitempty"`
	Border          *int32    `json:"border,omitempty"`
	DefaultAction   *string   `json:"defaultAction,omitempty"`
	Protocol        *string   `json:"protocol,omitempty"`
	IpList          []*string `json:"ipList,omitempty"`
}
