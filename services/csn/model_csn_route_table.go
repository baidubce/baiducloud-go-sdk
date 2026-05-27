package csn

type CsnRouteTable struct {
	CsnRtId     *string `json:"csnRtId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CsnType     *string `json:"type,omitempty"`
}
