package csn

type CsnRtRule struct {
	RuleId        *string `json:"ruleId,omitempty"`
	RouteType     *string `json:"routeType,omitempty"`
	CsnId         *string `json:"csnId,omitempty"`
	CsnRtId       *string `json:"csnRtId,omitempty"`
	Description   *string `json:"description,omitempty"`
	FromAttachId  *string `json:"fromAttachId,omitempty"`
	Status        *string `json:"status,omitempty"`
	SourceAddress *string `json:"sourceAddress,omitempty"`
	DestAddress   *string `json:"destAddress,omitempty"`
	NextHopId     *string `json:"nextHopId,omitempty"`
	NextHopName   *string `json:"nextHopName,omitempty"`
	NextHopRegion *string `json:"nextHopRegion,omitempty"`
	NextHopType   *string `json:"nextHopType,omitempty"`
	AsPath        *string `json:"asPath,omitempty"`
	Community     *string `json:"community,omitempty"`
	BlackHole     *bool   `json:"blackHole,omitempty"`
}
