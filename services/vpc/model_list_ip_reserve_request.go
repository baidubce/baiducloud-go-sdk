

package vpc

type ListIpReserveRequest struct {
    SubnetId *string `json:"-"`
    Marker *string `json:"-"`
    MaxKeys *int32 `json:"-"`
}