package rapidfs

type DescribeCceClustersRequest struct {
	VpcId   *string `json:"vpcId,omitempty"`
	MaxKeys *int32  `json:"maxKeys,omitempty"`
	Marker  *string `json:"marker,omitempty"`
}
