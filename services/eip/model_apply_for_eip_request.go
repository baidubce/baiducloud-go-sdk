

package eip

type ApplyForEipRequest struct {
    ClientToken *string `json:"-"`
    IpVersion *string `json:"ipVersion,omitempty"`
    RouteType *string `json:"routeType,omitempty"`
    BandwidthInMbps *int32 `json:"bandwidthInMbps,omitempty"`
    Billing *Billing `json:"billing,omitempty"`
    Name *string `json:"name,omitempty"`
    Tags []*TagModel `json:"tags,omitempty"`
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`
    AutoRenewTimeUnit *string `json:"autoRenewTimeUnit,omitempty"`
    AutoRenewTime *int32 `json:"autoRenewTime,omitempty"`
    DeleteProtect *bool `json:"deleteProtect,omitempty"`
}