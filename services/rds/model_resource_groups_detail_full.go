package rds

type ResourceGroupsDetailFull struct {
	Groups    []*ResourceInfo `json:"groups,omitempty"`
	Name      *string         `json:"name,omitempty"`
	RdsType   *string         `json:"type,omitempty"`
	Region    *string         `json:"region,omitempty"`
	Id        *string         `json:"id,omitempty"`
	Uuid      *string         `json:"uuid,omitempty"`
	Summary   *string         `json:"summary,omitempty"`
	Url       *string         `json:"url,omitempty"`
	AccountId *string         `json:"accountId,omitempty"`
	UserId    *string         `json:"userId,omitempty"`
	Tags      []*Tag          `json:"tags,omitempty"`
	ResUuid   *string         `json:"resUuid,omitempty"`
}
