package bcc

type DeploySetModel struct {
	DeploysetId           *string                    `json:"deploysetId,omitempty"`
	Name                  *string                    `json:"name,omitempty"`
	Desc                  *string                    `json:"desc,omitempty"`
	Strategy              *string                    `json:"strategy,omitempty"`
	Concurrency           *int32                     `json:"concurrency,omitempty"`
	AzIntstanceStatisList []*AzIntstanceStatisDetail `json:"azIntstanceStatisList,omitempty"`
}
