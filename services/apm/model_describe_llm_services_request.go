package apm

type DescribeLLMServicesRequest struct {
	BeginDatetime *string `json:"beginDatetime,omitempty"`
	EndDatetime   *string `json:"endDatetime,omitempty"`
	ServiceName   *string `json:"serviceName,omitempty"`
	ServiceId     *string `json:"serviceId,omitempty"`
	Env           *string `json:"env,omitempty"`
	Tag           *Tag    `json:"tag,omitempty"`
	OrderBy       *string `json:"orderBy,omitempty"`
	Order         *string `json:"order,omitempty"`
}
