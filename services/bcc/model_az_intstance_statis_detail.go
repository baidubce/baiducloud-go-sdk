package bcc

type AzIntstanceStatisDetail struct {
	ZoneName       *string   `json:"zoneName,omitempty"`
	InstanceCount  *int32    `json:"instanceCount,omitempty"`
	InstanceTotal  *int32    `json:"instanceTotal,omitempty"`
	BccInstanceCnt *int32    `json:"bccInstanceCnt,omitempty"`
	BbcInstanceCnt *int32    `json:"bbcInstanceCnt,omitempty"`
	InstanceIds    []*string `json:"instanceIds,omitempty"`
	BccInstanceIds []*string `json:"bccInstanceIds,omitempty"`
	BbcInstanceIds []*string `json:"bbcInstanceIds,omitempty"`
}
