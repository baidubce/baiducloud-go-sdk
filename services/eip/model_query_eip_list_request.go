

package eip

type QueryEipListRequest struct {
    IpVersion *string `json:"-"`
    Eip *string `json:"-"`
    InstanceType *string `json:"-"`
    InstanceId *string `json:"-"`
    Name *string `json:"-"`
    Status *string `json:"-"`
    EipIds []*string `json:"-"`
    Marker *string `json:"-"`
    MaxKeys *int32 `json:"-"`
}