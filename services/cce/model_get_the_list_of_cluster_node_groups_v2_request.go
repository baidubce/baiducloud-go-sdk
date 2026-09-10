package cce

type GetTheListOfClusterNodeGroupsV2Request struct {
	ClusterID         *string `json:"-"`
	PageNo            *int32  `json:"-"`
	PageSize          *int32  `json:"-"`
	KeywordType       *string `json:"-"`
	Keyword           *string `json:"-"`
	AutoscalerEnabled *string `json:"-"`
	ChargingType      *string `json:"-"`
}
