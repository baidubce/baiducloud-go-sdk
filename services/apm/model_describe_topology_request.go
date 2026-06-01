package apm

type DescribeTopologyRequest struct {
	Action        *string `json:"-"`
	ServiceName   *string `json:"serviceName,omitempty"`
	Env           *string `json:"env,omitempty"`
	BeginDatetime *string `json:"beginDatetime,omitempty"`
	EndDatetime   *string `json:"endDatetime,omitempty"`
}
