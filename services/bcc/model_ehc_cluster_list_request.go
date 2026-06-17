package bcc

type EhcClusterListRequest struct {
	EhcClusterIdList []*string `json:"ehcClusterIdList,omitempty"`
	NameList         []*string `json:"nameList,omitempty"`
	ZoneName         *string   `json:"zoneName,omitempty"`
}
