package scs

type NodeTypeItem struct {
	NodeType                *string  `json:"nodeType,omitempty"`
	InstanceFlavor          *int32   `json:"instanceFlavor,omitempty"`
	CpuNum                  *int32   `json:"cpuNum,omitempty"`
	NetworkThroughputInGbps *float32 `json:"networkThroughputInGbps,omitempty"`
	PeakQps                 *int32   `json:"peakQps,omitempty"`
	MaxConnections          *int32   `json:"maxConnections,omitempty"`
	AllowedNodeNumList      []*int32 `json:"allowedNodeNumList,omitempty"`
	MinDiskFlavor           *int32   `json:"minDiskFlavor,omitempty"`
	MaxDiskFlavor           *int32   `json:"maxDiskFlavor,omitempty"`
}
