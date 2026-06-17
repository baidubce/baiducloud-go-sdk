package bcc

type ListVolumeClustersRequest struct {
	Marker      *string `json:"-"`
	MaxKeys     *int32  `json:"-"`
	ZoneName    *string `json:"-"`
	ClusterName *string `json:"-"`
}
