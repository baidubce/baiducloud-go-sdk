package scs

type HotGroupSetFlowControlRulesRequest struct {
	GroupId       *string `json:"-"`
	ClusterShowId *string `json:"clusterShowId,omitempty"`
	QpsWrite      *int32  `json:"qpsWrite,omitempty"`
	QpsRead       *int32  `json:"qpsRead,omitempty"`
}
