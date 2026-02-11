

package eip

type EipBandwidthScalingCapacityRequest struct {
    Eip *string `json:"-"`
    ClientToken *string `json:"-"`
    NewBandwidthInMbps *int32 `json:"newBandwidthInMbps,omitempty"`
}